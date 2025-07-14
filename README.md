# Go + PostgreSQL CRUD PoC

This is a simple Go project demonstrating PostgreSQL integration with full CRUD operations and a REST API.

## Prerequisites
- Go 1.20+
- PostgreSQL running locally (default port 5432)

## Database Setup
Assumes a local PostgreSQL instance with:
- user: `postgres`
- password: `secret`
- database: `postgres`

You can start a local Postgres instance using Docker:
```sh
docker run --name pg-poc -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres
```

## Running the App
1. Install dependencies:
   ```sh
   go mod tidy
   ```
2. Run the server:
   ```sh
   go run .
   ```
   The server will start on `localhost:8080`.

## API Endpoints

### Create User
```
curl -X POST -H "Content-Type: application/json" -d '{"name":"Alice","email":"alice@example.com"}' http://localhost:8080/users
```

### Get All Users
```
curl http://localhost:8080/users
```

### Update User
```
curl -X PUT -H "Content-Type: application/json" -d '{"name":"Alice Updated","email":"alice2@example.com"}' http://localhost:8080/users/1
```

### Delete User
```
curl -X DELETE http://localhost:8080/users/1
```

## Project Structure
- `main.go` — Starts the server and handles HTTP routes
- `db.go` — Database connection and table creation
- `user.go` — User model and CRUD logic

## Notes
- All SQL queries use parameterized statements to prevent SQL injection.
- The users table is created automatically if it does not exist. 