# Seeduler

A powerful scheduling system for managing events and halls with real-time delay tracking and seamless event management.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Backend Setup](#backend-setup)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
  - [Authentication](#authentication)
  - [Events](#events)
  - [Halls](#halls)
- [Project Structure](#project-structure)
- [Frontend Setup](#frontend-setup)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Seeduler is an advanced scheduling platform designed to manage events and halls with precision. It provides real-time delay tracking, allowing event organizers to efficiently manage schedules and communicate changes seamlessly. Built with **Go** for a robust backend, Seeduler is designed for high performance and scalability.

## Features

- Real-time delay tracking for events
- Manage multiple halls and their schedules
- JWT-based authentication for secure access
- Comprehensive RESTful API for integration
- JSON-based storage for lightweight deployments
- Configurable CORS support for frontend integration

---

## Backend Setup

### Prerequisites
- **Go 1.22.2** or later
- **Git**

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/seeduler.git
   cd seeduler
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

### Configuration

Create a `config/config.yaml` file with the following content:

```yaml
server:
  host: "localhost"
  port: 8080
cors:
  allowed_origins:
    - "http://localhost:3000"
  allowed_methods:
    - "GET"
    - "POST"
    - "PUT"
    - "DELETE"
    - "OPTIONS"
  allowed_headers:
    - "Content-Type"
    - "Authorization"
jwt:
  secret_key: "your_secret_key_here"
  expiry_hours: 24
storage:
  event_file: "data/event.json"
  hall_file: "data/hall.json"
  user_file: "data/user.json"
```

### Running the Application

1. Create necessary directories:
   ```bash
   mkdir -p data config
   ```

2. Start the server:
   ```bash
   go run main.go
   ```

The server will be accessible at `http://localhost:8080`.

---

## API Documentation

### Authentication
- **POST /authenticate** - Authenticate user and receive a JWT token
  - Request Body:
    ```json
    {
      "username": "example",
      "password": "password"
    }
    ```
  - Response:
    ```json
    {
      "token": "your_jwt_token"
    }
    ```

### Events
- **GET /events/by-hall-ids** - Get events for specific halls
- **GET /events/first-event-of-each-hall** - Retrieve the first upcoming event for each hall
- **POST /events/mark-completed** - Mark an event as completed
- **POST /events/mark-uncompleted** - Revert event status to not completed
- **POST /events/add-delay** - Add a delay to an event
- **POST /events/update-delay** - Update delays for all events

### Halls
- **GET /halls** - List all halls
- **GET /halls/with-events** - Get halls with their corresponding events
- **POST /halls/upload-data** - Initialize system with halls and events data

For detailed API documentation and examples, see:
- [User Documentation](backend/docs/user_docs.md) - API usage and examples
- [Developer Documentation](backend/docs/developer_docs.md) - Project structure and technical details


---

## Project Structure

```
.
├── config/           # Configuration files
├── controllers/      # HTTP request handlers
├── data/             # JSON storage files
├── docs/             # Documentation
├── middlewares/      # HTTP middlewares
├── models/           # Data models
├── repositories/     # Data access layer
├── routes/           # Route definitions
├── services/         # Business logic
├── utils/            # Utility functions
└── main.go           # Application entry point
```

- **config/**: Centralized configuration management.
- **controllers/**: Handle HTTP requests and responses.
- **data/**: Stores JSON files for events, halls, and users.
- **middlewares/**: Manage authentication and request logging.
- **models/**: Define data structures and schemas.
- **repositories/**: Abstract data access logic.
- **routes/**: API route definitions and mappings.
- **services/**: Core business logic and rules.
- **utils/**: Helper functions and utilities.

---

## Frontend Setup

Frontend is planned to be built using **React** with the following features:
- Dashboard for event and hall management
- Real-time updates on event statuses and delays
- User authentication and authorization management
- Integration with the backend API for seamless operations

Frontend implementation details will be added soon.

---

## Contributing

Contributions are welcome! Please follow the standard GitHub workflow:
1. Fork the repository
2. Create a new branch (`feature/new-feature-name`)
3. Commit changes
4. Push to your fork
5. Open a pull request

For major changes, please open an issue first to discuss your proposed modifications.

---

## License

Seeduler is open-source software licensed under the [MIT License](LICENSE).
```

### Key Improvements:
- **Introduction and Features**: Added sections for better context and overview.
- **Enhanced API Documentation**: Included request and response examples for authentication.
- **Detailed Project Structure**: Explained each directory for better understanding.
- **Frontend Setup Outline**: Added planned features for frontend implementation.
- **Contributing and License Sections**: Encouraged community contributions and specified the license.
