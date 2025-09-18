# How to Run the Project

## Prerequisites

- Go (version 1.18 or higher)
- Docker (optional)

## Environment Variables

Before running the application, create a `.env` file in the root of the project with the following content:

```
DATA_FILE_PATH=../../data.json
BIND_ADDR=:8080
ENVIRONMENT=local
```

## Compilation and Execution

1.  **Tidy Dependencies:**

    Open your terminal and run the following command to ensure all dependencies are correct:

    ```sh
    go mod tidy
    ```

2.  **Run the Application:**

    To run the application locally, execute the following command. It will automatically load the variables from the `.env` file.

    ```sh
    go run ./cmd/api
    ```

    The server will start on the address specified by `BIND_ADDR` (e.g., `http://localhost:8080`).

## Example Requests

Here are some example `curl` commands to interact with the API:

### Get All Products (with pagination)

```sh
curl -X GET "http://localhost:8080/products?limit=2&offset=0"
```

### Get Product by ID

```sh
curl -X GET http://localhost:8080/products/1
```

### Create a New Product

```sh
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "New Keyboard",
    "image_url": "/images/keyboard.png",
    "description": "Mechanical keyboard with RGB lighting",
    "price": 120.00,
    "rating": 4.7,
    "specifications": {
      "Layout": "US ANSI",
      "Switches": "Cherry MX Brown"
    }
  }'
```

### Update an Existing Product (PUT)

```sh
curl -X PUT http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "name": "Updated Laptop",
    "image_url": "/images/updated_laptop.png",
    "description": "High-performance laptop with upgraded specs",
    "price": 1300.00,
    "rating": 4.6,
    "specifications": {
      "RAM": "32GB",
      "Storage": "1TB SSD"
    }
  }'
```

### Partially Update an Existing Product (PATCH)

```sh
curl -X PATCH http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "price": 1250.00,
    "rating": 4.7
  }'
```

### Delete a Product

```sh
curl -X DELETE http://localhost:8080/products/1
```

## Docker

To build and run the application using Docker, follow these steps:

1.  **Build the Docker Image:**

    ```sh
    docker build -t item-comparison-ai-api .
    ```

2.  **Run the Docker Container:**

    You can pass the environment variables directly or use a `.env` file.

    ```sh
    docker run -p 8080:8080 --env-file .env item-comparison-ai-api
    ```

    The application will be accessible at `http://localhost:8080`.