# Go Server

A simple HTTP server built with Go that serves static files and handles form submissions.

## Features

- Serves static files from `./static` directory
- `/hello` - GET endpoint that returns "hello!"
- `/form` - POST endpoint that processes form data (name and address)

## Requirements

- Go 1.25.0 or higher

## How to Run

1. Start the server:
   ```bash
   go run main.go
   ```

2. The server will start on port 8080. Open your browser and visit:
   - `http://localhost:8080` - Static files
   - `http://localhost:8080/hello` - Hello endpoint
   - `http://localhost:8080/form` - Form endpoint

## Endpoints

- `GET /` - Serves static files
- `GET /hello` - Returns "hello!"
- `POST /form` - Processes form data with `name` and `address` fields
