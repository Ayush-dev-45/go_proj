# Bookstore Management API

A REST API for managing books, built with Go, Gorilla Mux, and GORM with a MySQL backend.

## Tech Stack

- **Go** 1.25
- **Gorilla Mux** – HTTP router
- **GORM** – ORM for MySQL
- **MySQL** – Database

## Project Structure

```
bookstore-management/
├── cmd/
│   └── main/
│       └── main.go          # Application entry point, server on :9090
├── pkg/
│   ├── config/
│   │   └── app.go           # DB connection (MySQL DSN)
│   ├── controllers/
│   │   └── book.controllers.go  # HTTP handlers for book CRUD
│   ├── models/
│   │   └── book.go          # Book model and DB operations
│   ├── routes/
│   │   └── book.route.go    # Route definitions
│   └── utils/
│       └── utils.go        # Helpers (e.g. request body parsing)
├── go.mod
├── go.sum
└── README.md
```

## Prerequisites

- Go 1.25+
- MySQL (local or Docker)

## Database Setup

1. Create a MySQL database and user, e.g.:

   ```sql
   CREATE DATABASE mydatabase;
   CREATE USER 'myuser'@'%' IDENTIFIED BY 'mypassword';
   GRANT ALL ON mydatabase.* TO 'myuser'@'%';
   FLUSH PRIVILEGES;
   ```

2. Update the DSN in `pkg/config/app.go` if needed (user, password, host, port, database name). Default is:

   ```
   myuser:mypassword@tcp(127.0.0.1:3306)/mydatabase?charset=utf8mb4&parseTime=true&loc=Local
   ```

   The `Book` table is created automatically via GORM AutoMigrate on startup.

## Running the Application

From the **project root** (`bookstore-management/`):

```bash
go run ./cmd/main/main.go
```

Or from inside `cmd/main/`:

```bash
cd cmd/main
go run main.go
```

The server listens on **http://localhost:9090**.

## API Endpoints

| Method | Endpoint     | Description        |
|--------|--------------|--------------------|
| GET    | `/book/`     | List all books     |
| GET    | `/book/{id}` | Get book by ID     |
| POST   | `/book/`     | Create a book      |
| PUT    | `/book/{id}` | Update a book      |
| DELETE | `/book/{id}` | Delete a book      |

### Book Model

- `name` (string)
- `author` (string)
- `publication` (string)

The model also includes GORM’s `Model` (e.g. `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`).

### Examples

**Create a book (POST /book/)**

```bash
curl -X POST http://localhost:9090/book/ \
  -H "Content-Type: application/json" \
  -d '{"name":"The Go Programming Language","author":"Donovan & Kernighan","publication":"Addison-Wesley"}'
```

**Get all books (GET /book/)**

```bash
curl http://localhost:9090/book/
```

**Get book by ID (GET /book/1)**

```bash
curl http://localhost:9090/book/1
```

**Update a book (PUT /book/1)**

```bash
curl -X PUT http://localhost:9090/book/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Title","author":"New Author","publication":"New Pub"}'
```

**Delete a book (DELETE /book/1)**

```bash
curl -X DELETE http://localhost:9090/book/1
```

## License

MIT
