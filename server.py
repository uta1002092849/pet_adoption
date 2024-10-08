from concurrent import futures
import logging

import grpc
import uuid
import petAdoption_pb2
import petAdoption_pb2_grpc

from api.petAdoptionApi import PetAdoptionAPI

api = PetAdoptionAPI(None)

class PetAdoptionServicer(petAdoption_pb2_grpc.PetAdoptionServicer):
    
    def InsertPet(self, request, context):
        insert_pet = {
            'pet_name': request.name,
            'pet_gender': request.gender,
            'pet_age': request.age,
            'pet_breed': request.breed,
            'pet_picture': request.picture,
        }
        

        # Call the API
        response = api.insert_pet(**insert_pet)
        if response:
            return petAdoption_pb2.InsertPetResponse(success=True)
        else:
            return petAdoption_pb2.InsertPetResponse(success=False)
    
    def SearchPet(self, request, context):
        search_phrase = request.searchTerm
        result = api.search_pet(search_phrase)
        response = []
        for match in result:
            response.append(petAdoption_pb2.Pet(
                id= str(match.pet_id),
                name=match.pet_name,
                gender=match.pet_gender,
                age = match.age,
                breed = match.breed,
                picture = match.picture
            ))
        return petAdoption_pb2.SearchPetResponse(pets=response)


def serve():
    port = "50051"
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    petAdoption_pb2_grpc.add_PetAdoptionServicer_to_server(PetAdoptionServicer(), server)
    server.add_insecure_port("[::]:" + port)
    server.start()
    print("Server started, listening on " + port)
    server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig()
    serve()