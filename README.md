# GoLang-Gin
Implementing simple REST-API using Gin framework of golang

# Features
- Create a new book entry.
- Retrieve a list of all books
- Fetch a book by id.
- Issue a book by id .

# Prerequisites
- Go 1.19 installed on your machine
- Docker (optional).

# Installation
    1. Clone the repository to your local machine:
        git clone https://github.com/YashR461/GoLang-Gin.git

    2. Navigate to the project directory:
        cd GoLang-Gin
    3. Install the required Go packages:
        go mod download

# Directory Structure
GoLang-Gin
├── main.go
├── go.mod
├── go.sum
├── README.md


# Build & Run 
go mod tidy
Run the application using the following command: go run main.go
The application will start and listen on localhost:8080.

Use tools like  Postman to send HTTP requests to the provided endpoints (e.g., POST, PUT, PATCH, GET) to interact with the Books Management System.

# API Endpoints
1. POST /books – Creates a new entry for a book
2. GET /books – Displays the list of all books
3. GET /books/:id – Fetches the book by id
4. PATCH /books/issue/?: – Issues a book and reduces the quantity

| HTTP Method |                 Endpoint                   |      Description     |
|-------------|--------------------------------------------|----------------------|
| POST        | http://localhost:8080/books                | Adds a book          | 
| GET         | http://localhost:8080/books                | Get all books        |
| GET         | http://localhost:8080/books/id             | Update player by id  |                        
| PATCH       | http://localhost:8080/issue?1              | Get a random player  |                  
|-------------|------------------------------------------ -|----------------------|