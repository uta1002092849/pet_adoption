from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from models import Pet
from sqlalchemy import or_

# Constants for pet adoption database
DB_NAME = 'petadoptiondb'
SCHEMA_NAME = 'client_schema'
PET_TABLE_NAME = 'pet'
SQL_DB_IMPLEMENTATION = 'cockroachdb'
DEFAULT_HOST = 'localhost'
DEFAULT_PORT = '26257'

class PetAdoptionAPI:
    def __init__(self, db_uri):
        """Initialize the PetAdoptionAPI object."""
        if not db_uri:
            db_uri = f"{SQL_DB_IMPLEMENTATION}://root@{DEFAULT_HOST}:{DEFAULT_PORT}/{DB_NAME}?sslmode=disable"
        self.engine = create_engine(db_uri)
        self.Session = sessionmaker(bind=self.engine)
        self.session = self.Session()

    def insert_pet(self, pet_name, pet_gender, pet_age, pet_breed, pet_picture):
        """
        Insert a new pet into the database.
        Args:
            pet_name (str): The name of the pet.
            pet_gender (str): The gender of the pet.
            pet_age (int): The age of the pet.
            pet_breed (str): The breed of the pet.
            pet_picture (bytes): The picture of the pet.
        Returns:
            bool: True if the pet was successfully inserted, False otherwise.
        """
        new_pet = Pet(
            pet_name=pet_name,
            pet_gender=pet_gender,
            age=pet_age,
            breed=pet_breed,
            picture=pet_picture
        )
        
        try:
            self.session.add(new_pet)
        except Exception as e:
            print(f"Failed to insert a new pet: {e}")
            self.session.rollback()
            return False
        else:
            self.session.commit()
            return True
        
    
    def _try_parse_int(self, search_term):
        """
        Try to parse the search term as an integer.
        Args:
            search_term (str): The search term to parse.
        Returns:
            int: The integer value if the search term can be parsed as an integer, None otherwise.
        """
        try:
            return int(search_term)
        except ValueError:
            return None
    
    def search_pet(self, search_term):
        """
        Search for pets based on a key detail (name, gender, age, breed).
        Specifically, case-insensitive substring matching is used.
        Args:
            search_term (str): The search term to match against pet details.
        Returns:
            list[Pet]: A list of pets that match the search term.
        """
        # convert search term to substring format for string fields
        like_term = f"%{search_term}%"
        
        # try to parse the search term as an integer for age comparison
        age_term = self._try_parse_int(search_term)

        pets = self.session.query(Pet).filter(
            or_(
                Pet.pet_name.ilike(like_term),
                Pet.pet_gender.ilike(like_term),
                Pet.breed.ilike(like_term),
                Pet.age == age_term if age_term is not None else False
            )
        ).all()

        return pets
