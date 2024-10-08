CREATE DATABASE IF NOT EXISTS PetAdoptionDB;

USE PetAdoptionDB;

CREATE USER IF NOT EXISTS client;
GRANT CREATE ON DATABASE PetAdoptionDB TO client;

DROP SCHEMA IF EXISTS client_schema;
CREATE SCHEMA client_schema AUTHORIZATION client;

CREATE TABLE PetAdoptionDB.client_schema.Pet (
    id                  UUID NOT NULL DEFAULT gen_random_uuid(),
    name                STRING NOT NULL,
    gender              STRING NOT NULL,
    age                 OID NOT NULL,
    breed               STRING NOT NULL,
    picture             BYTES,
    PRIMARY KEY (id)
);