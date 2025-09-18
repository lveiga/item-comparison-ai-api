# Product Comparison API

This is a simple RESTful API built with Golang and the Gin framework to provide product details for a comparison feature.

## API Design

The API exposes the following endpoints for product management:

- `GET /products`: Returns a list of all products with optional pagination (`limit`, `offset`).
- `GET /products/{id}`: Returns details for a single product.
- `POST /products`: Creates a new product.
- `PUT /products/{id}`: Updates an existing product.
- `PATCH /products/{id}`: Partially updates an existing product.
- `DELETE /products/{id}`: Deletes a product.

### Product Model

The `Product` model includes the following fields:

- `id` (integer)
- `name` (string)
- `image_url` (string)
- `description` (string)
- `price` (float)
- `rating` (float)
- `category` (string)
- `specifications` (map[string]string)

## Setup and Running

To set up and run the project, please see the instructions in `run.md`.

## Running Unit Tests

To run the unit tests for the handlers, execute the following command:

```sh
go test ./internal/handlers
```

## Running Integration Tests

To run the integration tests, execute the following command from the root of the project:

```sh
go test -v ./...
```

## Architectural Decisions

- **Framework:** Gin was chosen for its lightweight nature and high performance.
- **Data Storage:** Product data is stored in an in JSON file to keep the project simple and avoid external dependencies like a database.
- **Structure:** The project is organized into `internal/` packages for a clean and scalable structure.

---

## Project Structure and How It Works

Below is a detailed description of the directory structure and the application's startup flow.

### `cmd/api/main.go`

This is the **entry point** of the application. The `main.go` file orchestrates the initialization of all essential components, including:
- The web framework (Gin).
- The general configurations (`config`).
- The `logger` system.
- The database connection (`database`).

After preparing the components, it initializes the server (`server`) and injects the necessary dependencies, such as routes and middlewares.

### `internal/server`

- **`app.go`**: Contains the core server logic. It is responsible for starting and stopping the application, checking the environment (development or production), and exposing the `healthcheck` endpoint.
- **`bindable.go`**: Performs the "bind" of all routes defined in the application to the Gin server, ensuring they are registered and ready to receive requests.

### `config`

This directory contains the general application settings, responsible for loading and managing environment variables required for the system to run.

### `internal/tests`

Stores the API's integration tests, which validate the complete application flows to ensure stability and the correct functioning of the endpoints.

### `internal/database`

Provides an abstraction layer for the database connection, centralizing its logic and facilitating future maintenance.

### `internal/handlers`

Contains the business logic for each endpoint. The handlers are responsible for receiving the HTTP request, processing it (by interacting with the repositories), and returning the response to the client.

### `internal/logger`

Abstracts the configuration of the logging system, allowing for centralized control over how application information is recorded.

### `internal/middleware`

Location for the API's middlewares. Currently, it contains the `healthcheck` middleware, which exposes an endpoint to check the application's health.

### `internal/models`

Defines the application's entities and data structures. At the moment, it only contains the `Product` model.

### `internal/repositories`

Implements the repository pattern to abstract data access. It contains a base implementation (`base_repository.go`) that can be extended and the `product_repository.go` specific to products.

### `internal/routes`

Defines all the API's endpoints. It currently contains the CRUD routes for the product resource, mapping each route to its respective handler.