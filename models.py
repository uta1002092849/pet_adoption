from sqlalchemy import Column, Integer, String, LargeBinary
from sqlalchemy.dialects.postgresql import OID
from sqlalchemy.orm import declarative_base

Base = declarative_base()

class Pet(Base):
    """The Pet class corresponds to the client_schema.Pet table."""
    __tablename__ = 'pet'
    __table_args__ = {'schema': 'client_schema'}  # Specify schema

    pet_id = Column(Integer, primary_key=True)
    pet_name = Column(String, nullable=False)
    pet_gender = Column(String, nullable=False)
    age = Column(OID, nullable=False)
    breed = Column(String, nullable=False)
    picture = Column(LargeBinary)

    def __repr__(self):
        """Return a human-readable representation of a Pet."""
        return f"Pet(pet_id={self.pet_id}, pet_name={self.pet_name}, pet_gender={self.pet_gender}, age={self.age}, breed={self.breed})"