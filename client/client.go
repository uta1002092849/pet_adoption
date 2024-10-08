package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "github.com/packer/petAdoption"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	// Set up a connection to the gRPC server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to the server with the following error code: %v", err)
	}
	defer conn.Close()
	client := pb.NewPetAdoptionClient(conn)

	// Set up HTTP server
	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/insert-pet", func(w http.ResponseWriter, r *http.Request) {
		handleInsertPet(w, r, client)
	})
	http.HandleFunc("/search-pet", func(w http.ResponseWriter, r *http.Request) {
		handleSearchPet(w, r, client)
	})

	fmt.Println("Server is running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handleInsertPet(w http.ResponseWriter, r *http.Request, client pb.PetAdoptionClient) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, "Unable to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Get form values
	name := r.FormValue("name")
	gender := r.FormValue("gender")
	ageStr := r.FormValue("age")
	breed := r.FormValue("breed")

	// Convert age to int32
	age, err := strconv.ParseInt(ageStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid age: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Get the picture file
	file, _, err := r.FormFile("picture")
	if err != nil {
		http.Error(w, "Error retrieving the file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the file content
	pictureData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading the file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	petRequest := &pb.InsertPetRequest{
		Name:    name,
		Gender:  gender,
		Age:     int32(age),
		Breed:   breed,
		Picture: pictureData,
	}

	response, err := client.InsertPet(ctx, petRequest)
	if err != nil {
		http.Error(w, "Could not insert pet: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"success": response.Success})
}

func handleSearchPet(w http.ResponseWriter, r *http.Request, client pb.PetAdoptionClient) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	searchTerm := r.URL.Query().Get("term")
	if searchTerm == "" {
		http.Error(w, "Search term is required", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SearchPet(ctx, &pb.SearchPetRequest{SearchTerm: searchTerm})
	if err != nil {
		http.Error(w, "Could not search for pets: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string][]*pb.Pet{"pets": response.Pets})
}
