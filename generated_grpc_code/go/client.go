package main

import (
	"context"
	"flag"
	"log"
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
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPetAdoptionClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Define a new pet to insert
	pet_request := &pb.InsertPetRequest{
		Name:   "Frodo",
		Gender: "Non-binary",
		Age:    100291,
		Breed:  "Fucking dog",
	}

	// Call the translate service
	r, err := c.InsertPet(ctx, pet_request)
	if err != nil {
		log.Fatalf("could not call insert pet: %v", err)
	}

	if r.Success {
		log.Printf("Pet inserted successfully")
	} else {
		log.Printf("Pet insertion failed")
	}

}
