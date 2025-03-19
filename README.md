# Task Management Backend

This is the backend for the Task Management application built using Golang, Gin, MongoDB, and WebSockets. The backend handles user authentication, task creation, assignment, and real-time notifications.

## Features

- User authentication (Register/Login)
- JWT-based authentication middleware
- Task creation and assignment
- Fetch tasks assigned to a user
- Mark tasks as completed
- Real-time notifications via WebSockets
- Deployment on Render

## Tech Stack

- **Backend:** Golang, Gin framework
- **Database:** MongoDB (MongoDB Atlas cluster)
- **Real-time communication:** WebSockets
- **Authentication:** JWT
- **Deployment:** Render

---

## Folder Structure

```
task_management/
│── backend/
│   ├── config/
│   │   ├── database.go
│   ├── controllers/
│   │   ├── authController.go
│   │   ├── taskController.go
│   ├── middleware/
│   │   ├── authMiddleware.go
│   ├── models/
│   │   ├── task.go
│   │   ├── userModel.go
│   ├── routes/
│   │   ├── authRoutes.go
│   │   ├── taskRoutes.go
│   ├── websockets/
│   │   ├── ws.go
│   ├── .env
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
```

---

## Installation & Setup

### Prerequisites

- Go installed (>= 1.18)
- MongoDB Atlas or local MongoDB setup
- Render account for deployment

### Steps to Run Locally

1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/task_management.git
   cd task_management/backend
   ```
2. Create a `.env` file in `backend/` and add:
   ```sh
   MONGO_URI="mongodb+srv://gshikhar84:ZlXQ1WC5tAe7Ka2C@cluster0.qmjcz.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
   JWT_SECRET="secret_key"
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Run the server:
   ```sh
   go run main.go
   ```

---

## API Endpoints & Postman Testing

### 1. Register a User

**Endpoint:** `POST /register`

**Request Body:**

```json
{
  "username": "shikhar_gupta",
  "password": "shikhar123"
}
```

**Response:**

```json
{
  "message": "Signup successful"
}
```

![alt text](Images/image.png)

### 2. Login

**Endpoint:** `POST /login`

**Request Body:**

```json
{
  "username": "shikhar_gupta",
  "password": "shikhar123"
}
```

**Response:**

```json
{
  "token": "your_jwt_token",
  "username": "testuser"
}
```

![alt text](Images/image-1.png)

### 3. Create a Task

**Endpoint:** `POST /tasks/create`

**Headers:**

```json
Authorization: Bearer <jwt_token>
```

**Request Body:**

```json
{
  "title": "task1",
  "description": "No Description",
  "assigned_to": "user2"
}
```

**Response:**

```json
{
  "id": "67daf3e42fe804a06f426041",
  "message": "Task created",
  "task": {
    "id": "67daf3e42fe804a06f426041",
    "title": "task1",
    "description": "No description",
    "assigned_to": "user2",
    "created_by": "shikhar_gupta",
    "status": "Pending",
    "created_at": 1742402532
  }
}
```

![alt text](Images/image-2.png)
![alt text](Images/image-3.png)

### 4. Get Assigned Tasks

**Endpoint:** `GET /tasks`

**Headers:**

```json
Authorization: Bearer <jwt_token>
```

**Response:**

```json
[
  {
    "id": "67daf3e42fe804a06f426041",
    "title": "task1",
    "description": "No description",
    "assigned_to": "user2",
    "created_by": "shikhar_gupta",
    "status": "Pending",
    "created_at": 1742402532
  }
]
```

![alt text](Images/image-4.png)

### 5. Mark Task as Done

**Endpoint:** `PUT /tasks/:id/done`

**Headers:**

```json
Authorization: Bearer <jwt_token>
```

**Response:**

```json
{
  "message": "Task marked as done"
}
```

![alt text](Images/image-5.png)
![alt text](Images/image-6.png)

### 6. Establish WebSocket Connection

**Endpoint:** `GET /ws?username=<your_username>`

**Description:**

- This establishes a WebSocket connection for real-time task notifications.
  ![alt text](Images/image-7.png)

---

## Deployment on Render

### Steps to Deploy

1. Initialize a Git repository (if not already):
   ```sh
   git init
   git add .
   git commit -m "Initial commit"
   ```
2. Push code to GitHub:
   ```sh
   git remote add origin https://github.com/gshikhar2021/task_management_backend.git
   git push -u origin main
   ```
3. Log in to [Render](https://dashboard.render.com/) and create a new **Web Service**.
4. Connect your GitHub repository.
5. Set up the `render.yaml` file:

   ```yaml
   services:
     - type: web
       runtime: go
       name: task-management-backend
       plan: free
       region: oregon
       buildCommand: |
         go mod tidy
         go build -o app .
       startCommand: ./app
       envVars:
         - key: PORT
           value: "8080"
         - key: MONGO_URI
           sync: false
         - key: JWT_SECRET
           sync: false
   ```

6. Add **MONGO_URI** and **JWT_SECRET** in Render’s environment variables.
7. Deploy the service and note the Render-provided URL.

Your backend is now live! 🎉

---

## License

This project is licensed under the MIT License.

---

## Author

Your Name - [GitHub](https://github.com/gshikhar2021)
