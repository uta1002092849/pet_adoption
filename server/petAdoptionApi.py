from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from sqlalchemy import or_
from sqlalchemy import Column, String, LargeBinary
from sqlalchemy.dialects.postgresql import OID
from sqlalchemy.orm import declarative_base
import uuid

Base = declarative_base()
# Constants for pet adoption database
DB_NAME = 'petadoptiondb'
SCHEMA_NAME = 'client_schema'
PET_TABLE_NAME = 'pet'
SQL_DB_IMPLEMENTATION = 'cockroachdb'
DEFAULT_HOST = 'roach1'
DEFAULT_PORT = '26257'

class Pet(Base):
    """The Pet class corresponds to the client_schema.Pet table."""
    __tablename__ = 'pet'
    __table_args__ = {'schema': 'client_schema'}  # Specify schema

    id = Column(String, primary_key=True)
    name = Column(String, nullable=False)
    gender = Column(String, nullable=False)
    age = Column(OID, nullable=False)
    breed = Column(String, nullable=False)
    picture = Column(LargeBinary)

    def __repr__(self):
        """Return a human-readable representation of a Pet."""
        return f"Pet(pet_id={self.id}, pet_name={self.name}, pet_gender={self.gender}, age={self.age}, breed={self.breed})"

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
        
        # Validate the input data
        if not pet_name or pet_gender.lower() not in ['male', 'female'] or pet_age <= 0 or not pet_breed or not pet_picture:
            return False
        
        new_pet = Pet(
            id = str(uuid.uuid4()),
            name=pet_name,
            gender=pet_gender,
            age=pet_age,
            breed=pet_breed,
            picture=pet_picture
        )
        
        # Check if the pet already exists
        existing_pet = self.search_pet(pet_name)
        # compare the pet details to check if the pet already exists
        if existing_pet:
            return False
            
        
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
                Pet.name.ilike(like_term),
                Pet.gender.ilike(like_term),
                Pet.breed.ilike(like_term),
                Pet.age == age_term if age_term is not None else False
            )
        ).all()
        return pets
