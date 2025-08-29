# Worlder Microservice Task

This project is a microservice-based system for producing and consuming sensor data. It consists of two main services: a `sensor-producer` that generates and publishes sensor data via MQTT, and a `sensor-consumer` that subscribes to the MQTT topic, processes the data, stores it in a database, and exposes a REST API for data retrieval and management.

## System Architecture

The system is composed of several microservices orchestrated by Docker Compose:

*   **Sensor Producers**: Three instances of the `sensor-producer` service are configured, each simulating a different sensor:
    *   `sensor-producer-temperature-1`: Simulates a temperature sensor.
    *   `sensor-producer-humidity`: Simulates a humidity sensor.
    *   `sensor-producer-temperature-2`: Simulates a second temperature sensor.
    Each producer generates data and publishes it to an MQTT topic. They also expose an HTTP endpoint to dynamically change their data generation frequency.

*   **`sensor-consumer`**: This service subscribes to the MQTT topic to receive data from all producers. It then stores this data in a MySQL database and provides a RESTful API for clients to authenticate and manage the sensor data.

*   **`nginx`**: A reverse proxy that sits in front of the `sensor-consumer`. All API requests to the consumer service are routed through Nginx.

For a visual representation of the architecture and database schema, please refer to the `System Diagram.jpg` and `ERD.jpg` files in the repository.

![System Diagram](System%20Diagram.jpg)

![ERD](ERD.jpg)

### Features

*   **Microservice Architecture**: Decoupled services for production and consumption of data.
*   **Asynchronous Communication**: Utilizes MQTT for efficient and reliable messaging between services.
*   **RESTful APIs**: Both services expose RESTful APIs:
    *   The `sensor-producer` API allows administrators to control the data generation frequency.
    *   The `sensor-consumer` API allows users to retrieve sensor data and administrators to manage it.
*   **Authentication & Authorization**: Secures endpoints using JWT and role-based access control (admin, user).
*   **Containerized**: Fully containerized using Docker for easy setup and deployment.

### Technologies Used

*   **Backend**: Go
*   **Web Framework**: Echo
*   **Database**: MySQL
*   **ORM**: GORM
*   **Messaging**: MQTT (Eclipse Paho)
*   **Containerization**: Docker, Docker Compose
*   **Authentication**: JSON Web Tokens (JWT)

## Getting Started

### Prerequisites

*   Docker
*   Docker Compose

### Installation

1.  **Clone the repository:**
    ```sh
    git clone <repository-url>
    cd worlder-microservice-task
    ```

2.  **Create a `.env` file:**
    Create a file named `.env` in the root of the project with the following variables.

    ```env
    # Application Ports
    APP_PORT=8080

    # Database Configuration (for sensor-consumer)
    DB_HOST=your-mysql-host
    DB_PORT=3306
    DB_USER=your-db-user
    DB_PASSWORD=your-db-password
    DB_NAME=your-db-name

    # MQTT Broker Configuration
    MQTT_HOST=your-mqtt-broker-host
    MQTT_PORT=1883
    MQTT_TOPIC=sensors/data
    MQTT_USERNAME=
    MQTT_PASSWORD=
    MQTT_PRODUCER_CLIENT_ID=sensor-producer

    # Security
    JWT_SECRET=your-jwt-secret-key
    DEFAULT_FREQUENCY=10
    ```

3.  **Run the application:**
    Use Docker Compose to build and run the services.
    ```sh
    docker-compose up -d --build
    ```

4.  **Access the services:**
    *   **Sensor Consumer API (via Nginx)**: `http://localhost:9081`
    *   **Sensor Producer APIs**:
        *   Temperature 1: `http://localhost:8080`
        *   Humidity: `http://localhost:8081`
        *   Temperature 2: `http://localhost:8082`

## Database Migration and Seeding

This project does not perform automatic database migrations or seeding upon startup. You must perform these steps manually.

### Migration

The database schema is defined by several GORM entities within the `sensor-consumer` service:
*   `SensorData`
*   `User`
*   `Role`
*   `UserRole`

To create the necessary tables, you can use GORM's `AutoMigrate` feature. You would typically do this by creating a temporary Go program or a CLI command that initializes a database connection and runs the migration.

Example of a migration script:
```go
package main

import (
    "fmt"
    "sensor-consumer/config"
    "sensor-consumer/core/entity"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func main() {
    // Assume you have a config loader
    // For simplicity, we'll manually define the DSN here
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    fmt.Println("Running Migrations...")
    db.AutoMigrate(&entity.SensorData{}, &entity.User{}, &entity.Role{}, &entity.UserRole{})
    fmt.Println("Migrations complete.")
}
```

### Seeding

After running the migrations, the database will be empty. You will need to manually add data, especially for `roles` and `users`, to be able to log in and use the API.

1.  **Create Roles**: You should add at least an `admin` and a `user` role to the `roles` table.
2.  **Create a User**: Create a new user in the `users` table. Remember to hash the password using a secure algorithm like bcrypt.
3.  **Assign Role to User**: Create an entry in the `user_roles` table to link the new user to a role.

## Postman Collection

You can find a Postman collection for this API here:

[https://www.postman.com/science-specialist-44124245/workspace/my-workspace/collection/31096617-69522de2-c912-4a04-8273-61a8acc8de02](https://www.postman.com/science-specialist-44124245/workspace/my-workspace/collection/31096617-69522de2-c912-4a04-8273-61a8acc8de02?action=share&creator=31096617&active-environment=31096617-fe682184-ad36-419e-81be-f36a5a7f5f53)

## API Endpoints

### Sensor Producer

The following endpoint is available on each of the three producer instances:
*   `http://localhost:8080` (Temperature 1)
*   `http://localhost:8081` (Humidity)
*   `http://localhost:8082` (Temperature 2)

*   **Change Sensor Frequency**
    *   **Endpoint**: `POST /sensor/:frequency`
    *   **Description**: Changes the frequency (in seconds) at which the specific sensor instance publishes data.
    *   **Authorization**: `Admin` role required.
    *   **Path Parameters**:
        *   `frequency` (integer, **required**): The interval in milliseconds for publishing sensor data.
    *   **Example**: `POST http://localhost:8080/sensor/5`

### Sensor Consumer (`http://localhost:9081`)

#### Authentication

*   **Login**
    *   **Endpoint**: `POST /login`
    *   **Description**: Authenticates a user and returns a JWT token.
    *   **Request Body**:
        ```json
        {
            "username": "your_username",
            "password": "your_password"
        }
        ```

#### Sensor Data

*   **Get Sensor Data**
    *   **Endpoint**: `GET /sensor`
    *   **Description**: Retrieves a paginated and filterable list of sensor data.
    *   **Authorization**: `User` or `Admin` role required.
    *   **Query Parameters**:
        *   `page` (integer, **required**): The page number for pagination.
        *   `length` (integer, **required**): The number of items per page.
        *   `sort` (string, **required**): The field to sort by (e.g., `created_at`).
        *   `order` (string, **required**): The sort order (`asc` or `desc`).
        *   `id1` (string, optional): Filter by a specific sensor type name (e.g., `Temperature`).
        *   `id2` (integer, optional): Filter by a specific sensor ID.
        *   `time_start` (datetime, optional): The start of a time range filter (Format: `YYYY-MM-DDTHH:MM:SSZ`).
        *   `time_end` (datetime, optional): The end of a time range filter (Format: `YYYY-MM-DDTHH:MM:SSZ`).

*   **Update Sensor Data**
    *   **Endpoint**: `PATCH /sensor` or `PATCH /sensor/:id`
    *   **Description**: Updates one or more sensor data entries. Can be used to update a single entry by its ID in the path, or multiple entries by using query parameters as filters.
    *   **Authorization**: `Admin` role required.
    *   **Path Parameters**:
        *   `id` (integer, optional): The ID of a specific sensor entry to update.
    *   **Query Parameters (for bulk updates)**:
        *   `id1`, `id2`, `time_start`, `time_end`: Used to filter the records to be updated if no `:id` is provided.
    *   **Request Body**:
        ```json
        {
            "sensor_value": 99.9
        }
        ```

*   **Delete Sensor Data**
    *   **Endpoint**: `DELETE /sensor` or `DELETE /sensor/:id`
    *   **Description**: Deletes one or more sensor data entries. Can be used to delete a single entry by its ID in the path, or multiple entries by using query parameters as filters.
    *   **Authorization**: `Admin` role required.
    *   **Path Parameters**:
        *   `id` (integer, optional): The ID of a specific sensor entry to delete.
    *   **Query Parameters (for bulk deletes)**:
        *   `id1`, `id2`, `time_start`, `time_end`: Used to filter the records to be deleted if no `:id` is provided.
