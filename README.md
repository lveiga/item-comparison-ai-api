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