
# User Documentation

## Introduction

This project provides an API for managing events and halls. It includes endpoints for user authentication, managing events, and managing halls.

## Getting Started

### Prerequisites

- Go 1.22.2 or later
- A terminal or command prompt

### Configuration

Before running the application, set up the configuration file at `config/config.json` with the JWT secret key:

```json
{
  "jwt_secret_key": "your_secret_key"
}
```

### Running the Application

1. Clone the repository:
   ```sh
   git clone https://github.com/your-repo/seeduler.git
   cd seeduler
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Run the application:
   ```sh
   go run main.go
   ```

The server will start at `http://localhost:8080`.

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

## Conclusion

This documentation provides an overview of the API endpoints and how to use them. For more detailed information, refer to the developer documentation.