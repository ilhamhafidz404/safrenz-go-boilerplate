# 🚀 Safrenz Go Boilerplate

Starter template Backend Safrenz

---

## Quick Start

```bash
# 1. Clone & install dependencies
git clone [https://github.com/username/safrenz-go-boilerplate.git](https://github.com/username/safrenz-go-boilerplate.git)
cd safrenz-go-boilerplate
go mod tidy

# 2. Setup Environment
cp .env.example .env

# 3. Run Application
go run main.go

```

- Base URL: http://localhost:3000
- Swagger UI: http://localhost:3000/swagger/index.html

---

## Folder Structure

* **`config/`**: Initializes environment variables, DB connection, and Redis.
* **`internal/bootstrap/`**: Assembles (wires up) all dependencies between layers.
* **`internal/dto/`**: Contains request validation schemas and JSON response formats.
* **`internal/handler/`**: Handles HTTP requests/responses and input validation.
* **`internal/model/`**: Defines database table entity structs (GORM).
* **`internal/repository/`**: Isolates database query operations (CRUD).
* **`internal/service/`**: Handles core business logic and data aggregation (Third-Party/Redis).
* **`pkg/platform/`**: Houses the Redis client and Third-Party API HTTP clients.
* **`pkg/utils/`**: Utility helpers for JSON response formatting & image upload validation (magic bytes).
* **`router/`**: Registers API endpoints, Swagger UI, and static file serving.