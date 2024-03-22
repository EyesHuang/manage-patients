# Manage Patients

This project facilitates interaction with patient data through two primary web services.

## Available Services

The project offers two main services. Detailed API documentation is located in the `backend/doc` directory, under the `swagger.yaml` file. To explore and interact with the API documentation, please use [Swagger Editor](https://editor.swagger.io/).

- **GET /patients**: Retrieves a list of all patients.
- **PUT /orders/{id}**: Updates an existing medical order by its unique identifier.

## Getting Started

Below are the instructions for setting up and running the backend and frontend services of the project.

### Prerequisites

Before starting, ensure you have the following installed:
- Docker and Docker Compose for database setup.
- Go, for running the backend services.
- Node.js and npm, for running the frontend application.

### Setting Up PostgreSQL

1. Make sure Docker is running on your system.
2. Use the following command to start the PostgreSQL server using Docker Compose:

    ```shell
    docker-compose up -d
    ```

This command initializes and starts a PostgreSQL instance as defined in the `docker-compose.yml` file.

### Starting the Backend Services

To get the backend services up and running:

1. Open a terminal and navigate to the `backend/cmd/patient` directory.
2. Execute the Go program with:

    ```shell
    cd backend/cmd/patient
    go run main.go
    ```

This starts the backend web services, making them ready to accept requests.

### Starting the Frontend Application

To run the frontend application:

1. In the project's root directory, install the necessary npm packages (if it's the first time running the application):

    ```shell
    npm install
    ```

2. Start the application with:

    ```shell
    npm start
    ```

The frontend application will now be running and accessible in your web browser, typically at `http://localhost:3000`.
