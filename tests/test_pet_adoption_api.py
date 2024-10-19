import pytest
from server.petAdoptionApi import PetAdoptionAPI

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
        'pet_picture': b'fake_pet_picture_data'
    }
    
    # Act
    result = pet_api.insert_pet(**pet_data)
    
    # Assert
    assert result is True
    
    # Verify the pet was actually inserted
    inserted_pet = pet_api.search_pet('TestDog')[0]
    assert inserted_pet.name == pet_data['pet_name']
    assert inserted_pet.gender == pet_data['pet_gender']
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
        'pet_picture': b'fake_cat_pet_picture'
    }
    pet_api.insert_pet(**test_pet)
    
    # Act
    result = pet_api.search_pet('SearchTest')
    
    # Assert
    assert len(result) > 0
    found_pet = result[0]
    assert found_pet.name == test_pet['pet_name']
    assert found_pet.gender == test_pet['pet_gender']
    assert found_pet.age == test_pet['pet_age']
    assert found_pet.breed == test_pet['pet_breed']

    # Clean up: remove the test pet from the database
    pet_api.session.delete(found_pet)
    pet_api.session.commit()

def test_insert_duplicate_pet(pet_api):
    # Arrange: Insert a test pet
    pet_data = {
        'pet_name': 'DuplicateTestDog',
        'pet_gender': 'Male',
        'pet_age': 2,
        'pet_breed': 'Labrador',
        'pet_picture': b'fake_pet_picture_data'
    }
    pet_api.insert_pet(**pet_data)

    # Act: Try to insert the same pet again
    result = pet_api.insert_pet(**pet_data)

    # Assert: Verify that the second insertion fails
    assert result is False  # Assuming the API returns False on duplicate insertion

    # Clean up: remove the test pet from the database
    inserted_pet = pet_api.search_pet('DuplicateTestDog')[0]
    pet_api.session.delete(inserted_pet)
    pet_api.session.commit()

def test_insert_pet_invalid_data(pet_api):
    # Arrange: Prepare invalid pet data (missing required fields, etc.)
    invalid_pet_data = {
        'pet_name': '',  # Invalid: pet_name is empty
        'pet_gender': 'Male',
        'pet_age': -1,  # Invalid: pet_age cannot be negative
        'pet_breed': 'Labrador',
        'pet_picture': b'fake_pet_picture_data'
    }

    # Act: Attempt to insert invalid pet data
    result = pet_api.insert_pet(**invalid_pet_data)

    # Assert: Verify that the insertion fails
    assert result is False  # Assuming the API returns False on invalid data

def test_search_pet_by_substring(pet_api):
    # Arrange: Insert multiple test pets
    pets_to_insert = [
        {
            'pet_name': 'FluffyCat',
            'pet_gender': 'Female',
            'pet_age': 2,
            'pet_breed': 'Persian',
            'pet_picture': b'fake_fluffy_picture'
        },
        {
            'pet_name': 'FluffyDog',
            'pet_gender': 'Male',
            'pet_age': 3,
            'pet_breed': 'Golden Retriever',
            'pet_picture': b'fake_fluffy_dog_picture'
        },
        {
            'pet_name': 'SparkyDog',
            'pet_gender': 'Male',
            'pet_age': 1,
            'pet_breed': 'Beagle',
            'pet_picture': b'fake_sparky_picture'
        }
    ]
    
    for pet in pets_to_insert:
        pet_api.insert_pet(**pet)

    # Act: Search for pets with the substring "Fluffy"
    result = pet_api.search_pet('Fluffy')

    # Assert: Verify that the search returns the correct pets
    assert len(result) == 2  # Should find 2 pets with the name containing "Fluffy"
    
    # Verify the details of the found pets
    found_pets_names = [pet.name for pet in result]
    assert 'FluffyCat' in found_pets_names
    assert 'FluffyDog' in found_pets_names
    assert 'SparkyDog' not in found_pets_names  # Ensure unrelated pet is not included

    # Clean up: remove the test pets from the database
    for pet in result:
        pet_api.session.delete(pet)
    pet_api.session.commit()