import pytest
from api.petAdoptionApi import PetAdoptionAPI

# Constants for database connection (adjust these as needed)
DB_URI = "cockroachdb://root@localhost:26257/petadoptiondb?sslmode=disable"

@pytest.fixture
def pet_api():
    """Create a PetAdoptionAPI instance connected to the real database."""
    api = PetAdoptionAPI(DB_URI)
    yield api
    # Clean up (if necessary) after the test
    api.session.close()

def test_insert_pet_success(pet_api):
    # Arrange
    pet_data = {
        'pet_name': 'TestDog',
        'pet_gender': 'Male',
        'pet_age': 2,
        'pet_breed': 'Labrador',
        'pet_picture': b'fake_picture_data'
    }
    
    # Act
    result = pet_api.insert_pet(**pet_data)
    
    # Assert
    assert result is True
    
    # Verify the pet was actually inserted
    inserted_pet = pet_api.search_pet('TestDog')[0]
    assert inserted_pet.pet_name == pet_data['pet_name']
    assert inserted_pet.pet_gender == pet_data['pet_gender']
    assert inserted_pet.age == pet_data['pet_age']
    assert inserted_pet.breed == pet_data['pet_breed']
    assert inserted_pet.picture == pet_data['pet_picture']

    # Clean up: remove the test pet from the database
    pet_api.session.delete(inserted_pet)
    pet_api.session.commit()

def test_search_pet(pet_api):
    # Arrange: Insert a test pet
    test_pet = {
        'pet_name': 'SearchTestCat',
        'pet_gender': 'Female',
        'pet_age': 3,
        'pet_breed': 'Siamese',
        'pet_picture': b'fake_cat_picture'
    }
    pet_api.insert_pet(**test_pet)
    
    # Act
    result = pet_api.search_pet('SearchTest')
    
    # Assert
    assert len(result) > 0
    found_pet = result[0]
    assert found_pet.pet_name == test_pet['pet_name']
    assert found_pet.pet_gender == test_pet['pet_gender']
    assert found_pet.age == test_pet['pet_age']
    assert found_pet.breed == test_pet['pet_breed']

    # Clean up: remove the test pet from the database
    pet_api.session.delete(found_pet)
    pet_api.session.commit()