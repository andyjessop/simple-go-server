# Simple Go HTTP Server

This project is a simple HTTP server built in Go (Golang), demonstrating basic CRUD operations without the use of external frameworks. The server allows the creation, retrieval, and deletion of posts, which are simple data structures stored in memory. This project is intended for educational purposes, showcasing Go's capabilities in handling HTTP requests, managing concurrency, and structuring a simple REST API.

There is an accompanying blog post, which you can find here: [Building a Basic HTTP Server in Go: A Step by Step Tutorial](https://dev.to/andyjessop/building-a-basic-http-server-in-go-a-step-by-step-tutorial-ma4).

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing.

### Prerequisites

Before running the server, ensure you have Go installed on your machine. The project was built using Go version 1.22.0, but it should be compatible with most Go versions. You can download Go from [the official website](https://go.dev/learn/).

### Installing

To set up the project, clone the repository to your local machine:

```bash
git clone https://github.com/yourusername/simple-go-server.git
cd simple-go-server
```

Then, run the server using:

```bash
go run main.go
```

The server will start listening on http://localhost:8080. You can interact with the API through this base URL.

### Running the tests

To execute the automated tests for this system, run the following command in the project directory:

```bash
go test
```

### API Endpoints

The server defines the following RESTful endpoints:

```
GET /posts: Retrieve all posts.
GET /posts/{id}: Retrieve a single post by its ID.
POST /posts: Create a new post. The request body should contain the body field.
DELETE /posts/{id}: Delete a post by its ID.
```
