CREATE DATABASE IF NOT EXISTS PetAdoptionDB;

USE PetAdoptionDB;

CREATE USER IF NOT EXISTS client;
GRANT CREATE ON DATABASE PetAdoptionDB TO client;

DROP SCHEMA IF EXISTS client_schema;
CREATE SCHEMA client_schema AUTHORIZATION client;

CREATE TABLE PetAdoptionDB.client_schema.Pet (
    pet_id              UUID NOT NULL DEFAULT gen_random_uuid(),
    pet_name            STRING NOT NULL,
    pet_gender          STRING NOT NULL,
    age                 OID NOT NULL,
    breed               STRING NOT NULL,
    picture             BYTES,
    PRIMARY KEY (pet_id)
);