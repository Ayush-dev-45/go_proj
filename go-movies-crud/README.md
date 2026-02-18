# Go Movies CRUD

A simple REST API for managing a movie catalog. Built with Go and the [Gorilla Mux](https://github.com/gorilla/mux) router, it provides full CRUD (Create, Read, Update, Delete) operations for movies.

## What It Does

This API stores movies in memory with the following data:
- **ID** – Unique identifier
- **ISBN** – Book/movie catalog number
- **Title** – Movie title
- **Director** – First and last name

The server comes preloaded with two sample movies and listens on port `8080`.

## API Endpoints

| Method | Endpoint        | Description           |
|--------|-----------------|-----------------------|
| GET    | `/movies`       | List all movies       |
| GET    | `/movies/{id}`  | Get a movie by ID     |
| POST   | `/movies`       | Create a new movie    |
| PUT    | `/movies/{id}`  | Update a movie by ID  |
| DELETE | `/movies/{id}`  | Delete a movie by ID  |

## Prerequisites

- [Go 1.25+](https://go.dev/dl/)

## How to Run

1. **Clone or navigate to the project**
   ```bash
   cd go-movies-crud
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Start the server**
   ```bash
   go run main.go
   ```

   You should see:
   ```
   Server started at port 8080
   ```

4. **Build a binary (optional)**
   ```bash
   go build -o go-movies-crud
   ./go-movies-crud
   ```

## Usage Examples

**List all movies**
```bash
curl http://localhost:8080/movies
```

**Get a movie by ID**
```bash
curl http://localhost:8080/movies/1
```

**Create a movie**
```bash
curl -X POST http://localhost:8080/movies \
  -H "Content-Type: application/json" \
  -d '{"isbn":"123456","title":"Inception","director":{"firstname":"Christopher","lastname":"Nolan"}}'
```

**Update a movie**
```bash
curl -X PUT http://localhost:8080/movies/1 \
  -H "Content-Type: application/json" \
  -d '{"isbn":"438227","title":"Movie One Updated","director":{"firstname":"John","lastname":"Doe"}}'
```

**Delete a movie**
```bash
curl -X DELETE http://localhost:8080/movies/1
```
