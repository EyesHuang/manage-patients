# manage-patients

This project primarily provides two web services for interacting with patient data.

## Available Services
This project offers two key services. Detailed API documentation can be found within the `doc` directory in a file named `api.yaml`. To view and interact with the API documentation, open the file using [Swagger Editor](https://editor.swagger.io/).

- `GET /patients`: Retrieve information for all patients.
- `PUT /orders/{id}`: Update a specified medical order using its ID.

## Getting Started

### Setting Up PostgreSQL
Start the PostgreSQL server using Docker Compose:

    ```shell
    docker-compose up -d
    ```

### Starting the Web Services

To start the web services, navigate to the `cmd/patient` directory and run the Go program:

```shell
cd cmd/patient
go run main.go