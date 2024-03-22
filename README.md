# Simple CRUD App in Golang

This project aims to provide basic CRUD operations for managing posts via a RESTful API using Go. It demonstrates how to connect to a MySQL database, define data models, handle HTTP requests, and interact with the database using raw SQL queries.

## Installation

- Install Go Using this site

```bash
  https://go.dev/doc/install
```

- Clone Repository

```bash
  https://github.com/DrCodeNinja/CRUD-App-Golang.git
```

- Open project using any code editor
- Create .env file in your project root directory.
- Place this code with your credentials.

```bash
  PORT=3000
  DB_URL="<username>:<password>@tcp(localhost:3306)/<database name>"
```

- Below code for running database migration scripts

```bash
  go run migrate/migrate.go
```

- Automatically recompile and restart the Go application during development

```bash
  CompileDaemon -command="./first"
```

- You can test this application with postman

#### Enjoy....
