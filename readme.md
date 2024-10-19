# Pet Adoption Project

Welcome to the Pet Adoption Project! This project aims to help connect potential pet owners with pets in need of a home.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)

## Introduction

The Pet Adoption Project is a simple web application backed by CockroachDB distributed SQL database and gRPC. It allows users to insert and search for available pets in the database. The goal of this project is to learn gRPC and Docker.

## Features

- Insert pets into the database with name, gender, age, breed, and picture.
- Search for pets using a search phrase that matches any of the fields.

## Installation

To get started with the project, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/mmt2849/pet_adoption.git
   ```
2. Navigate to the project directory:
   ```bash
   cd pet_adoption
   ```
3. Fetch CockroachDB images and start the cluster (you might need to use `sudo` for all subsequent steps):
   ```bash
   make db-up
   ```
   This will spin up 3 distributed containers joined together in a Docker private network, each with its own CPU core (requires at least 3 cores for database alone).
4. Perform one-time initialization for the cluster database:
   ```bash
   make db-init
   ```
5. Perform one-time table initialization for the database:
   ```bash
   make db-table-init
   ```
   This will initialize an empty table with the schema defined in the `init.sql` file. Optionally, you can inspect the table at various points during the program runtime using `make db-table-show`.
6. Build and run the client and server services:
   ```bash
   make service-up
   ```
   This will attempt to build both client and server images using the respective Docker files in the `root/client` and `root/server` subfolders. After this step is done, everything is ready, and the web interface can be accessed at `localhost:8081`. This is the only accessible port from the host machine. All other services like the server and database host are hidden behind the private network.

## Usage

Navigate to `http://localhost:8081` to start using the app. You can specify pet information, and the picture is required (tested with .jpg format). Once some pets have been inserted, you can search for them by name, breed, age, etc. As long as your search phrase is a substring of any of these fields, the matching pets will be returned and rendered with their respective pictures.
