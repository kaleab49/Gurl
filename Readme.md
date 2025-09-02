# Gurl

**Gurl** is a simple clone of the popular `curl` CLI tool, implemented in **Golang**.  
It allows you to send HTTP requests directly from the command line with support for multiple methods, JSON payloads, and verbose output.  

---

## ✨ Features

- **Supports HTTP Methods**  
  - `GET` (default)  
  - `POST`  
  - `PUT`  
  - `DELETE`  

- **Send Data**  
  - Use `-data` flag to send raw JSON in POST/PUT requests.  
  - Example:  
    ```bash
    go run main.go -url https://example.com/resource -request POST -data '{"key":"value"}'
    ```

- **Automatic JSON Header**  
  - Automatically sets `Content-Type: application/json` when `-data` is provided for POST/PUT requests.  

- **Verbose Mode**  
  - Use `-verbose` to see detailed request info including method, URL, headers, and body.  
  - Example:  
    ```bash
    go run main.go -url https://example.com/resource -request POST -data '{"key":"value"}' -verbose
    ```

- **Simple CLI Flags**  
  - `-url` → specify the URL to request  
  - `-request` → specify HTTP method (GET/POST/PUT/DELETE)  
  - `-data` → raw JSON data to send in the request body  
  - `-verbose` → print detailed request info  

- **Cross-platform**: Works anywhere Go runs (Linux, macOS, Windows).  
- **Lightweight**: No external dependencies—just Go standard library.  

---

## 🚀 Usage Examples

```bash
# Simple GET request
go run main.go -url https://jsonplaceholder.typicode.com/posts

# POST request with JSON
go run main.go -url https://jsonplaceholder.typicode.com/posts -request POST -data '{"title":"Hello","body":"World","userId":1}'

# PUT request
go run main.go -url https://jsonplaceholder.typicode.com/posts/1 -request PUT -data '{"title":"Updated Title","body":"Updated Body","userId":1}'

# DELETE request
go run main.go -url https://jsonplaceholder.typicode.com/posts/1 -request DELETE

# Verbose mode to see full request details
go run main.go -url https://jsonplaceholder.typicode.com/posts -request POST -data '{"title":"Test","body":"Body","userId":1}' -verbose
