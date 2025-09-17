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

    Execute the following command to start the API server:

    ```sh
    go run ./cmd/api
    ```

    The server will start on `http://localhost:8080`.

## Example Request

To retrieve a product, send a GET request to the `/products/{id}` endpoint. For example:

```sh
cURL -X GET http://localhost:8080/products/1
```