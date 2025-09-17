# Product Comparison API

This is a simple RESTful API built with Golang and the Gin framework to provide product details for a comparison feature.

## API Design

The API exposes a single endpoint to retrieve product information:

- `GET /products/{id}`: Returns details for a single product.

### Product Model

The `Product` model includes the following fields:

- `id` (integer)
- `name` (string)
- `image_url` (string)
- `description` (string)
- `price` (float)
- `rating` (float)
- `specifications` (map[string]string)

## Setup and Running

To set up and run the project, please see the instructions in `run.md`.

## Running Unit Tests

To run the unit tests for the handlers, execute the following command:

```sh
go test ./handlers
```

## Running Integration Tests

To run the integration tests, execute the following command from the root of the project:

```sh
go test -v ./...
```

## Architectural Decisions

- **Framework:** Gin was chosen for its lightweight nature and high performance.
- **Data Storage:** Product data is stored in an in-memory slice to keep the project simple and avoid external dependencies like a database.
- **Structure:** The project is organized into `handlers` and `models` packages for a clean and scalable structure.
