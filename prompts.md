# Initial prompt to create boilerplate

```
# Start with a clear and concise objective for the AI.
> I need to build a simple RESTful API in Golang using the Gin framework. The API will provide product details for a comparison feature.

# Define the technical requirements and constraints.
>
> 1.  **Objective:** Create a backend API for a product comparison feature.
> 2.  **API Endpoint:** Design and implement a RESTful GET endpoint at `/products/{id}` to return details for a single product. The response should include:
>     * `id` (integer)
>     * `name` (string)
>     * `image_url` (string)
>     * `description` (string)
>     * `price` (float)
>     * `rating` (float)
>     * `specifications` (map[string]string)
>
> 3.  **Data:** Use a simple JSON file or an in-memory struct slice to store product data. Do not use a database for this project.
> 4.  **Error Handling:** Implement basic error handling for scenarios like "product not found" (returning a 404 status code) and invalid requests.
> 5.  **Code Structure:** Organize the project with a clean, logical file structure (e.g., `main.go`, `models/`, `handlers/`).
> 6.  **Non-functional Requirements:** Include inline comments to explain the logic, especially for handlers and data structures.
> 7.  **Documentation:** Generate a `README.md` file that explains the API design, setup instructions, and key architectural decisions.
> 8.  **Project Files:** Create a `run.md` with instructions on how to compile and run the project.

# Request the AI to generate the code and documentation.
>
> **Task Breakdown:**
> 1.  Generate the initial project structure, including the `main.go` file.
> 2.  Create a Go struct for the `Product` model with JSON tags.
> 3.  Generate the handler function for the `/products/{id}` endpoint, including the logic to handle the GET request and error scenarios.
> 4.  Provide example JSON data for at least three products.
> 5.  Generate the `README.md` and `run.md` files based on the project requirements.
> 6.  Ensure all generated code follows Golang best practices and includes inline comments.
>
> **Final Output:** Provide all the necessary code and documentation in separate code blocks. Do not include any conversational text, just the files and their content.
```


# Prompt to Unit Test
## This prompt is focused on testing internal logic and functions in isolation, without the need to run the server.
```
> I need to create unit tests for my Golang API. The project uses the Gin framework. I want to test the `GetProductByID` handler function, which retrieves a product from an in-memory map based on an ID.

>
> **Task Breakdown:**
> 1. **Test Case 1: Success Scenario.**
>    * Simulate a GET request with a valid product ID (e.g., `1`).
>    * Assert that the HTTP status code is `200 OK`.
>    * Assert that the JSON response body matches the expected product data.
>
> 2. **Test Case 2: Product Not Found.**
>    * Simulate a GET request with an invalid product ID that does not exist (e.g., `999`).
>    * Assert that the HTTP status code is `404 Not Found`.
>    * Assert that the JSON response body contains an appropriate error message.
>
> 3. **Test Case 3: Invalid ID Format.**
>    * Simulate a GET request with an invalid ID format (e.g., a string like `"abc"`).
>    * Assert that the HTTP status code is `400 Bad Request`.
>    * Assert that the JSON response body contains a message indicating the invalid ID.
>
> 4. **Code Structure:**
>    * Use the `net/http/httptest` package to create a mock HTTP server and recorder.
>    * Organize the tests in a separate file (e.g., `handlers_test.go`).
>    * Include inline comments to explain the purpose of each test case.
>
> **Final Output:** Generate the complete Go code for the unit tests, ready to be added to the project.
``` 

# Prompt to Integration Test
## This prompt focuses on testing the complete API flow, ensuring that the server and endpoints work together correctly.

```
> I need to create integration tests for my Golang API built with the Gin framework. The tests should cover the full API endpoint `/products/{id}`.

>
> **Task Breakdown:**
> 1. **Test Case 1: Valid API Call.**
>    * Start a real HTTP server in a separate goroutine.
>    * Use `net/http` to make an actual GET request to the `/products/1` endpoint.
>    * Assert that the HTTP status code is `200 OK`.
>    * Verify that the JSON response body contains the correct product details.
>
> 2. **Test Case 2: Endpoint Not Found.**
>    * Make a GET request to a non-existent endpoint (e.g., `/api/v1/products/1`).
>    * Assert that the HTTP status code is `404 Not Found`.
>    * Verify that the response body indicates the route was not found.
>
> 3. **Test Setup and Teardown:**
>    * Use a testing `main` function or a test suite to set up and tear down the server.
>    * Use `t.Parallel()` to allow tests to run in parallel.
>
> 4. **Code Structure:**
>    * Organize the tests in a separate file (e.g., `integration_test.go`).
>    * Include inline comments to explain the test logic.
>
> **Final Output:** Generate the complete Go code for the integration tests.
```



# Prompt to Restructure Go Project Folders
## This prompt instructs Gemini to reorganize the project so that it follows Go best practices, separating internal logic, the application entry point, and integration tests.

```
> I need to refactor my Golang API project to follow a more scalable and idiomatic project layout.

> **Current Project Structure:**
> ```
> /item-comparison-ai-api
> ├── main.go
> ├── handlers/
> ├── models/
> ├── integration_test.go
> ```

> **Target Structure:**
> Refactor the project to follow the standard Go project layout.
>
> 1.  **`cmd/` folder:** Move the `main.go` file. The `main.go` file should be the only entry point for the application.
> 2.  **`internal/ folder:** Move folders handlers and models into internal/. and create a folder for file integration_test.go
```