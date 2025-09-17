# How to Run the Project

## Prerequisites

- Go (version 1.18 or higher)

## Compilation and Execution

1.  **Tidy Dependencies:**

    Open your terminal and run the following command to ensure all dependencies are correct:

    ```sh
    go mod tidy
    ```

2.  **Run the Application:**

    To run the application locally, you need to set the `DATA_FILE_PATH` environment variable.

    ```sh
    DATA_FILE_PATH=data.json go run ./cmd/api
    ```

    The server will start on `http://localhost:8080`.

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

    The `DATA_FILE_PATH` environment variable is set within the Dockerfile to `/root/data.json`.

    ```sh
    docker run -p 8080:8080 item-comparison-ai-api
    ```

    The application will be accessible at `http://localhost:8080`.
