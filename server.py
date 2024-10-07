from concurrent import futures
import logging

import grpc
from generated_grpc_code.python import petAdoption_pb2_grpc
from generated_grpc_code.python import petAdoption_pb2

from api.petAdoptionApi import PetAdoptionApi

api = PetAdoptionApi()

class Translator(petAdoption_pb2_grpc.PetAdoptionServicer):
    
    def InsertPet(self, request, context):
        insert_pet = {
            'pet_name': request.pet_name,
            'pet_gender': request.pet_gender,
            'pet_age': request.pet_age,
            'pet_breed': request.pet_breed,
            'pet_picture': request.pet_picture,
        }
        

        # Call the API
        response = api.InsertPet(**insert_pet)
        if response:
            return petAdoption_pb2.InsertPetResponse(success=True)
        else:
            return petAdoption_pb2.InsertPetResponse(success=False)
    
    def SearchPet(self, request, context):
        search_phrase = request.search_term
        result = api.SearchPet(search_phrase)
        response = []
        for match in result:
            response.append(petAdoption_pb2.Pet(
                pet_id= match.pet_id,
                pet_name=match.pet_name,
                pet_gender=match.pet_gender,
                age = match.age,
                breed = match.breed,
                picture = match.picture
            ))
        return petAdoption_pb2.SearchPetResponse(pets=response)


def serve():
    port = "50051"
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    petAdoption_pb2_grpc.add_TranslateServicer_to_server(Translator(), server)
    server.add_insecure_port("[::]:" + port)
    server.start()
    print("Server started, listening on " + port)
    server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig()
    serve()