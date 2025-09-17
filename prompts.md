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