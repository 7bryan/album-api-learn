# Go Gin RESTful Album API (MySQL Edition)

A simple, layered architecture RESTful API built using **Go (Golang)**, the **Gin Web Framework**, and **GORM** connected to a **MySQL database**.

## Project Architecture

This project follows a structured layout separation to isolate responsibilities cleanly:

- `config/`: Database setup & connectivity management.
- `controllers/`: Direct handlers managing requests and outputs.
- `models/`: Database schema structures and definitions.

## Project Structure

```
gin-album-api/
├── config/
│   └── database.go    # Database connection logic
├── controllers/
│   └── album.go       # HTTP Handlers (GET, POST logic)
├── models/
│   └── album.go       # Database schema definition
├── .env               # Environment variables (DB credentials)
├── .gitignore
├── go.mod
├── go.sum
├── main.go            # Entry point & routing setups
└── README.md
```

## API Specification

| Method | Endpoint  | Description                  | Request Body        |
| :----- | :-------- | :--------------------------- | :------------------ |
| `GET`  | `/albums` | Retrieve list of all records | None                |
| `POST` | `/albums` | Insert a new item into MySQL | JSON (Album object) |

### JSON Schema Blueprint

> **Note:** The `id` property auto-increments automatically in the database.

```json
{
  "title": "Abbey Road",
  "artist": "The Beatles",
  "price": 29.99
}
```

## Prerequisites & Setup

1. **Local System Tools**: Ensure you have installed **Go** and a **MySQL Server** instance.
2. **Setup Database**: Log in to MySQL client and initialize a target empty schema:
   ```sql
   CREATE DATABASE album_db;
   ```
3. **Environment Setup**: Clone the project, copy configurations, and populate `.env` variables matching your environment credentials:

   ```bash
   cp .env.example .env # or create manually
   ```

4. **Boot Strategy**:
   ```bash
   go mod download
   go run main.go
   ```

## Quick Test Hooks

### Fetch Albums (`GET`)

```bash
curl http://localhost:8080/albums
```

### Insert Album (`POST`)

```bash
curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"title": "Kind of Blue","artist": "Miles Davis","price": 24.50}'
```
