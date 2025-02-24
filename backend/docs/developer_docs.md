# Developer Documentation

## Project Structure

```
.
├── .env
├── .gitignore
├── config/
│   └── config.json
├── controllers/
│   ├── event_controller.go
│   ├── hall_controllers.go
│   └── user_controller.go
├── data/
│   ├── event.json
│   ├── hall.json
│   └── user.json
├── docs/
│   └── developer_docs.md
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── middlewares/
│   └── auth_middleware.go
├── models/
│   ├── Event.go
│   ├── Hall.go
│   ├── request_response.go
│   └── User.go
├── repositories/
│   ├── event_repository.go
│   ├── hall_repository.go
│   └── user_repository.go
├── routes/
│   ├── event_routes.go
│   ├── hall_routes.go
│   ├── routes.go
│   └── user_routes.go
├── services/
│   ├── event_service.go
│   ├── hall_Service.go
│   └── user_service.go
├── tests/
├── utils/
│   └── config.go
└── vendor/
```

## Configuration

The configuration file is located at `config/config.json`. It contains the JWT secret key used for authentication.

```json
{
    "jwt_secret_key": "your_secret_key"
}
```

## API Endpoints

### Authentication

- **POST /authenticate**
    - Request Body:
        ```json
        {
            "username": "Main Hall_user",
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

- **GET /events/by-hall-ids**
    - Query Parameters: `hall_ids` (comma-separated list of hall IDs)
    - Example:
        ```sh
        curl -X GET "http://localhost:8080/events/by-hall-ids?hall_ids=1,2,3"
        ```

- **GET /events/first-event-of-each-hall**
    - Example:
        ```sh
        curl -X GET "http://localhost:8080/events/first-event-of-each-hall"
        ```

- **POST /events/mark-completed**
    - Request Body:
        ```json
        {
            "event_id": 1
        }
        ```
    - Example:
        ```sh
        curl -X POST "http://localhost:8080/events/mark-completed" -H "Content-Type: application/json" -d '{"event_id": 1}'
        ```

- **POST /events/mark-uncompleted**
    - Request Body:
        ```json
        {
            "event_id": 1
        }
        ```
    - Example:
        ```sh
        curl -X POST "http://localhost:8080/events/mark-uncompleted" -H "Content-Type: application/json" -d '{"event_id": 1}'
        ```

- **POST /events/add-delay**
    - Request Body:
        ```json
        {
            "event_id": 1,
            "delay": "10m"
        }
        ```
    - Example:
        ```sh
        curl -X POST "http://localhost:8080/events/add-delay" -H "Content-Type: application/json" -d '{"event_id": 1, "delay": "10m"}'
        ```

- **POST /events/update-delay**
    - Example:
        ```sh
        curl -X POST "http://localhost:8080/events/update-delay"
        ```

### Halls

- **GET /halls**
    - Example:
        ```sh
        curl -X GET "http://localhost:8080/halls"
        ```

- **GET /halls/with-events**
    - Example:
        ```sh
        curl -X GET "http://localhost:8080/halls/with-events"
        ```

- **POST /halls/upload-data**
    - Request Body:
        ```json
        {
            "halls": [
                {
                    "id": 1,
                    "name": "Main Hall",
                    "delayed_time": 15,
                    "info": {}
                },
                {
                    "id": 2,
                    "name": "Conference Room",
                    "delayed_time": 10,
                    "info": {}
                },
                {
                    "id": 3,
                    "name": "Auditorium",
                    "delayed_time": 20,
                    "info": {}
                }
            ],
            "events": [
                {
                    "id": 1,
                    "title": "Event 1",
                    "hall_id": 1,
                    "scheduled_start_time": "2025-02-23T03:00:00Z",
                    "scheduled_end_time": "2025-02-23T05:00:00Z",
                    "start_time": "2025-02-23T03:00:00Z",
                    "end_time": "2025-02-23T05:00:00Z",
                    "is_completed": false,
                    "is_started": false,
                    "info": {}
                },
                {
                    "id": 2,
                    "title": "Event 2",
                    "hall_id": 2,
                    "scheduled_start_time": "2025-02-23T03:30:00Z",
                    "scheduled_end_time": "2025-02-23T05:30:00Z",
                    "start_time": "2025-02-23T03:30:00Z",
                    "end_time": "2025-02-23T05:30:00Z",
                    "is_completed": false,
                    "is_started": false,
                    "info": {}
                }
            ]
        }
        ```
    - Example:
        ```sh
        curl -X POST "http://localhost:8080/halls/upload-data" -H "Content-Type: application/json" -d '{
            "halls": [
                {
                    "id": 1,
                    "name": "Main Hall",
                    "delayed_time": 15,
                    "info": {}
                },
                {
                    "id": 2,
                    "name": "Conference Room",
                    "delayed_time": 10,
                    "info": {}
                },
                {
                    "id": 3,
                    "name": "Auditorium",
                    "delayed_time": 20,
                    "info": {}
                }
            ],
            "events": [
                {
                    "id": 1,
                    "title": "Event 1",
                    "hall_id": 1,
                    "scheduled_start_time": "2025-02-23T03:00:00Z",
                    "scheduled_end_time": "2025-02-23T05:00:00Z",
                    "start_time": "2025-02-23T03:00:00Z",
                    "end_time": "2025-02-23T05:00:00Z",
                    "is_completed": false,
                    "is_started": false,
                    "info": {}
                },
                {
                    "id": 2,
                    "title": "Event 2",
                    "hall_id": 2,
                    "scheduled_start_time": "2025-02-23T03:30:00Z",
                    "scheduled_end_time": "2025-02-23T05:30:00Z",
                    "start_time": "2025-02-23T03:30:00Z",
                    "end_time": "2025-02-23T05:30:00Z",
                    "is_completed": false,
                    "is_started": false,
                    "info": {}
                }
            ]
        }'
        ```

## Architecture

### Controllers

- `event_controller.go`: Handles event-related API endpoints.
- `hall_controllers.go`: Handles hall-related API endpoints.
- `user_controller.go`: Handles user authentication API endpoints.

### Services

- `event_service.go`: Contains business logic for events.
- `hall_Service.go`: Contains business logic for halls.
- `user_service.go`: Contains business logic for user authentication and JWT token generation.

### Repositories

- `event_repository.go`: Handles data persistence for events.
- `hall_repository.go`: Handles data persistence for halls.
- `user_repository.go`: Handles data persistence for users.

### Middlewares

- `auth_middleware.go`: Middleware for JWT authentication.

### Models

- `Event.go`: Defines the `Event` model.
- `Hall.go`: Defines the `Hall` model.
- `User.go`: Defines the `User` model.
- `request_response.go`: Defines request and response models.

### Utilities

- `config.go`: Utility functions for loading configuration.

### Data

- `event.json`: Stores event data.
- `hall.json`: Stores hall data.
- `user.json`: Stores user data.

### Configuration

- `config.json`: Stores configuration settings, including the JWT secret key.

### Main Entry Point

- `main.go`: The main entry point of the application. It initializes the services, controllers, and routes, and starts the HTTP server.

## Running the Project

1. Set up the configuration file at `config/config.json` with the JWT secret key.
2. Run the application:
     ```sh
     go run main.go
     ```
```
