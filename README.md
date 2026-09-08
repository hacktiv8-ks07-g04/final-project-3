# Hacktiv8 Go Group 4 GLNG-KS07 Final Project 3

Task Management API built with Go, Gin, GORM, and PostgreSQL.

## Deployment

> https://final-project-3-production-efec.up.railway.app/

## Admin Credentials

- Email: `admin@gmail.com`
- Password: `admin123`

> The admin user is seeded automatically on first startup when the users table is empty.

---

## Setup

### Prerequisites

- Go 1.26+
- PostgreSQL

### 1. Create the database

```sql
CREATE DATABASE final_project_3;
```

### 2. Create `.env` in the project root

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=final_project_3
JWT_SECRET_KEY=some-long-random-secret
PORT=8080
```

### 3. Run the application

```bash
go run main.go
```

Tables are auto-migrated and the admin user is seeded on first startup.

Alternatively, use [Air](https://github.com/air-verse/air) for live-reload:

```bash
air
```

---

## API Endpoints

All protected endpoints require an `Authorization: Bearer <token>` header.

### Error Response Format

```json
{
  "message": "error description",
  "status": 400,
  "error": "BAD_REQUEST"
}
```

Error codes: `BAD_REQUEST` (400), `NOT_AUTHENTICATED` (401), `NOT_AUTHORIZED` (403), `NOT_FOUND` (404), `INVALID_REQUEST_BODY` (422), `INTERNAL_SERVER_ERROR` (500).

---

### Users

| Method | Path                  | Auth | Description         |
|--------|-----------------------|------|---------------------|
| POST   | /users/register       | No   | Register a new user |
| POST   | /users/login          | No   | Login               |
| PUT    | /users/update-account | Yes  | Update own account  |
| DELETE | /users/delete-account | Yes  | Delete own account  |

### Categories

| Method | Path                    | Auth | Admin | Description      |
|--------|-------------------------|------|-------|------------------|
| GET    | /categories             | Yes  | No    | Get all categories |
| POST   | /categories             | Yes  | Yes   | Create category  |
| PATCH  | /categories/:categoryId | Yes  | Yes   | Update category  |
| DELETE | /categories/:categoryId | Yes  | Yes   | Delete category  |

### Tasks

| Method | Path                               | Auth | Description           |
|--------|------------------------------------|------|-----------------------|
| POST   | /tasks                             | Yes  | Create a task         |
| GET    | /tasks                             | Yes  | Get all tasks         |
| PUT    | /tasks/:taskId                     | Yes  | Update task title/desc |
| PATCH  | /tasks/update-status/:taskId       | Yes  | Update task status    |
| PATCH  | /tasks/update-category/:taskId     | Yes  | Update task category  |
| DELETE | /tasks/:taskId                     | Yes  | Delete a task         |

---

## Examples

### Register

```bash
curl -X POST http://localhost:8080/users/register \
  -H "Content-Type: application/json" \
  -d '{"full_name": "John Doe", "email": "john@example.com", "password": "secret123"}'
```

Response:

```json
{
  "status": 201,
  "message": "Successfully registered new user",
  "data": {
    "id": 1,
    "full_name": "John Doe",
    "email": "john@example.com",
    "created_at": "2026-09-08T12:00:00Z"
  }
}
```

### Login

```bash
curl -X POST http://localhost:8080/users/login \
  -H "Content-Type: application/json" \
  -d '{"email": "john@example.com", "password": "secret123"}'
```

Response:

```json
{
  "status": 200,
  "message": "successfully logged in",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

### Update Account

```bash
curl -X PUT http://localhost:8080/users/update-account \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"full_name": "John Smith", "email": "john@example.com"}'
```

Response:

```json
{
  "status": 200,
  "message": "successfully updated user",
  "data": {
    "id": 1,
    "full_name": "John Smith",
    "email": "john@example.com",
    "updated_at": "2026-09-08T12:00:00Z"
  }
}
```

### Create a Category (admin only)

```bash
curl -X POST http://localhost:8080/categories \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <admin_token>" \
  -d '{"type": "very-urgent"}'
```

Response:

```json
{
  "status": 200,
  "message": "Successfully created new category",
  "data": {
    "id": 1,
    "type": "very-urgent",
    "created_at": "2026-09-08T12:00:00Z"
  }
}
```

### Get All Categories

```bash
curl http://localhost:8080/categories \
  -H "Authorization: Bearer <token>"
```

Response:

```json
{
  "status": 200,
  "message": "Successfully get all categories",
  "data": [
    {
      "id": 1,
      "type": "very-urgent",
      "created_at": "2026-09-08T12:00:00Z",
      "updated_at": "2026-09-08T12:00:00Z",
      "Tasks": [
        {
          "id": 1,
          "title": "My Task",
          "status": false,
          "description": "Do something",
          "user_id": 1,
          "category_id": 1,
          "created_at": "2026-09-08T12:00:00Z",
          "updated_at": "2026-09-08T12:00:00Z"
        }
      ]
    }
  ]
}
```

### Update a Category (admin only)

```bash
curl -X PATCH http://localhost:8080/categories/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <admin_token>" \
  -d '{"type": "medium"}'
```

Response:

```json
{
  "status": 200,
  "message": "Successfully updated category",
  "data": {
    "id": 1,
    "type": "medium",
    "updated_at": "2026-09-08T12:00:00Z"
  }
}
```

### Delete a Category (admin only)

```bash
curl -X DELETE http://localhost:8080/categories/1 \
  -H "Authorization: Bearer <admin_token>"
```

Response:

```json
{
  "status": 200,
  "message": "Successfully deleted category"
}
```

### Create a Task

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"title": "My Task", "description": "Do something", "category_id": 1}'
```

Response:

```json
{
  "status": 201,
  "message": "Success create new task",
  "data": {
    "id": 1,
    "title": "My Task",
    "description": "Do something",
    "status": false,
    "user_id": 1,
    "category_id": 1,
    "created_at": "2026-09-08T12:00:00Z"
  }
}
```

### Get All Tasks

```bash
curl http://localhost:8080/tasks \
  -H "Authorization: Bearer <token>"
```

Response:

```json
{
  "status": 200,
  "message": "Success get all tasks",
  "data": [
    {
      "id": 1,
      "title": "My Task",
      "status": false,
      "description": "Do something",
      "user_id": 1,
      "category_id": 1,
      "created_at": "2026-09-08T12:00:00Z",
      "User": {
        "id": 1,
        "email": "john@example.com",
        "full_name": "John Doe"
      }
    }
  ]
}
```

### Update Task

```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"title": "My Updated Task", "description": "Do something else"}'
```

Response:

```json
{
  "status": 200,
  "message": "Success update task title and description field",
  "data": {
    "id": 1,
    "title": "My Updated Task",
    "description": "Do something else",
    "status": false,
    "user_id": 1,
    "category_id": 1,
    "updated_at": "2026-09-08T12:00:00Z"
  }
}
```

### Update Task Status

```bash
curl -X PATCH http://localhost:8080/tasks/update-status/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"status": true}'
```

Response:

```json
{
  "status": 200,
  "message": "Success update task status field",
  "data": {
    "id": 1,
    "title": "My Task",
    "description": "Do something",
    "status": true,
    "user_id": 1,
    "category_id": 1,
    "updated_at": "2026-09-08T12:00:00Z"
  }
}
```

### Delete a Task

```bash
curl -X DELETE http://localhost:8080/tasks/1 \
  -H "Authorization: Bearer <token>"
```

Response:

```json
{
  "status": 200,
  "message": "Task has been successfully deleted"
}
```

### Update Task Category

```bash
curl -X PATCH http://localhost:8080/tasks/update-category/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"category_id": 2}'
```

Response:

```json
{
  "status": 200,
  "message": "Success update task category field",
  "data": {
    "id": 1,
    "title": "My Task",
    "description": "Do something",
    "status": false,
    "user_id": 1,
    "category_id": 2,
    "updated_at": "2026-09-08T12:00:00Z"
  }
}
```

---

### Thank You
