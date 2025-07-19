# GO Project Overview

## Project Summary

This project is a RESTful API built with Go, designed for user authentication and management. It provides endpoints for user registration, login, profile management, and related operations. The application follows a modular structure with clear separation of concerns, making it easy to maintain and extend.

---

## Setup Instructions

### Prerequisites
- Go (version 1.18 or higher recommended)
- A running instance of a supported database (e.g., PostgreSQL, MySQL, or SQLite)

### Installation
1. **Clone the repository:**
   ```sh
   git clone <your-repo-url>
   cd GO
   ```
2. **Install dependencies:**
   ```sh
   go mod tidy
   ```
3. **Configure environment variables:**
   - Create a `.env` file or set environment variables as required (e.g., database connection string, port, JWT secret).
   - Example variables:
     ```env
     DB_HOST=localhost
     DB_PORT=5432
     DB_USER=youruser
     DB_PASSWORD=yourpassword
     DB_NAME=yourdb
     JWT_SECRET=your_jwt_secret
     PORT=8080
     ```

### Running the Application
- To start the server, run:
  ```sh
  go run main.go
  ```
- The API will be available at `http://localhost:8080` (or the port you specified).

---

## Project Structure

- **go.mod / go.sum**  
  These files manage your Go module’s dependencies.
  - `go.mod` declares the module path and required dependencies.
  - `go.sum` records checksums for dependency verification.

- **main.go**  
  The entry point of your application.
  - Sets up the server, loads configuration, connects to the database, and starts the HTTP server.
  - Imports routes, middleware, and other core components.

---

## 2. api/

This directory contains the core logic for handling HTTP requests.

### - controllers/

- **auth.controller.go**  
  Handles authentication-related endpoints (login, register, logout, etc.).
- **users.controller.go**  
  Manages user-related endpoints (fetching user info, updating profiles, etc.).

### - routes/

- **routes.go**  
  Defines all the HTTP routes/endpoints and maps them to their respective controller functions.
  - Example: `POST /login` → `auth.controller.LoginHandler`

---

## 3. database/

- **db.go**  
  Handles database connection and setup.
  - Contains functions to initialize the DB, manage migrations, and provide a DB instance to other parts of the app.

---

## 4. helper/

Utility functions and helpers used throughout the project.

- **common.go**  
  General-purpose helper functions (e.g., string manipulation, token generation).
- **config.go**  
  Loads and manages configuration (environment variables, config files).
- **resHandler.go**  
  Standardizes API responses (success, error formatting).

---

## 5. middleware/

- **jsonValidator.middleware.go**  
  Middleware for validating incoming JSON requests.
  - Ensures requests have the correct structure before reaching controllers.

---

## 6. model/

- **user.go**  
  Defines the User model/struct and possibly related methods.
  - Maps to the users table in the database.
  - May include methods for password hashing, validation, etc.

---

## 7. validations/

- **auth.validation.go**  
  Contains validation logic for authentication requests.
  - Example: Ensures login requests have both email and password, registration fields are valid, etc.

---

# How It All Fits Together

1. **main.go** starts the app, loads config, connects to the DB, and sets up routes.
2. **routes.go** defines endpoints and attaches middleware (like JSON validation).
3. **Controllers** handle the business logic for each endpoint.
4. **Models** represent and interact with database entities.
5. **Helpers** provide utility functions used across the app.
6. **Middleware** processes requests before they reach controllers (e.g., validating JSON).
7. **Validations** ensure incoming data is correct and safe.

---

# Typical Request Flow

1. **Client** sends HTTP request.
2. **Route** matches the request and applies middleware.
3. **Middleware** (e.g., JSON validator) checks the request.
4. **Controller** processes the request, using **models** and **helpers** as needed.
5. **Database** operations are performed via the **database/db.go** connection.
6. **Response** is formatted using helpers and sent back to the client.

---

If you want a deeper dive into any specific file or want to see code samples, let me know! 