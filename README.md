# Gopportunities

![Go Version](https://img.shields.io/badge/go-1.19-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)

An API for managing job opportunities, developed in Go.

## 📝 Project Description

Gopportunities is a CRUD (Create, Read, Update, Delete) system for job opportunities. The API allows listing, creating, updating, and deleting openings, serving as an excellent hands-on project for learning the fundamentals of Go, creating REST APIs, and integrating with databases.

## ✨ Features

-   [x] Create new job openings
-   [x] List all registered openings
-   [x] View a specific job opening
-   [x] Update a job opening's data
-   [x] Delete a job opening

## 🛠️ Technologies Used

This project was built using the following technologies:

-   **[Go](https://golang.org/)**: Main programming language.
-   **[Gin Gonic](https://gin-gonic.com/)**: HTTP framework for building the API.
-   **[SQLite](https://www.sqlite.org/index.html)**: Relational database for storing openings.
-   **[GORM](https://gorm.io/)**: ORM (Object-Relational Mapping) to interact with the database.
-   **[Swagger](https://swagger.io/)**: For interactive API documentation.

## 🚀 Getting Started

Follow the instructions below to run the project in your local environment.

### Prerequisites

You will need to have [Go](https://go.dev/doc/install) (version 1.19 or higher) installed on your machine.

### Installation

1.  Clone the repository:
    ```bash
    git clone [https://github.com/reykalencar/gopportunities.git](https://github.com/reykalencar/gopportunities.git)
    ```

2.  Navigate to the project directory:
    ```bash
    cd gopportunities
    ```

3.  Install the dependencies:
    ```bash
    go mod tidy
    ```

4.  Run the application:
    ```bash
    go run main.go
    ```

The API will be running at `http://localhost:8080`.

## 📖 API Usage

The complete and interactive API documentation is available via Swagger. After starting the server, access:

**[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**

### Main Endpoints

| Method | Endpoint          | Description                 |
| :----- | :---------------- | :-------------------------- |
| `POST` | `/api/v1/opening` | Creates a new opening.      |
| `GET`  | `/api/v1/opening` | Fetches a specific opening. |
| `GET`  | `/api/v1/openings`| Lists all openings.         |
| `PUT`  | `/api/v1/opening` | Updates an opening.         |
| `DELETE`| `/api/v1/opening` | Deletes an opening.         |
