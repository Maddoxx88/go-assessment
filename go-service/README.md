# Go PDF Report Service

This Go microservice logs in to a Node.js backend at startup, retrieves authentication cookies (accessToken, refreshToken, csrfToken), and uses them to handle secure PDF report generation.

---

## 🚀 Features

- Automatic login to backend on startup
- Stores and uses cookies for authenticated requests
- Compatible with Go 1.21+
- Designed for integration with school management backend

---

## 📦 Prerequisites

Before running this service, make sure you have:

1. **Go** (v1.21 or later)  
   [Download Go](https://go.dev/dl/)

2. **Node.js backend** running locally or remotely  
   - Must expose:
     - `POST /api/v1/auth/login` – Login endpoint
     - `GET /api/v1/students/:id` – Fetch student details
   - Must allow cookie-based authentication for local development

3. **Environment variables** for backend connection:
   - Backend API URL
   - Backend login credentials (email & password)

---

## ⚙️ Setup

1. **Clone the repository**

   ```bash
   git clone https://github.com/Maddoxx88/go-assessment.git
   cd go-assessment/go-service

2. **Create .env file**

   ```bash
   git clone https://github.com/<your-username>/<your-repo>.git
   cd <your-repo>/go-service

3. **Install dependencies**

   ```bash
   go mod tidy

4. **Run the service**

   ```bash
   go run main.go router.go

If successful you'll see:

✅ Logged in and tokens stored at startup

🚀 Go service started on :8080

📁 Project Structure

```
go-service/
├── go.sum            # Go dependencies lock file
├── go.mod            # Go module file
├── main.go           # Main entry point for the service
├── router.go            # Registers endpoints (/ping, /report)
├── handlers/                # HTTP handlers (controllers) 
│   ├── report_handler.go        # Handles report generation
├── models/           
│   ├── student.go            # Student Model
├── services/           
│   ├── report_service.go            # Fetches student data from Node.js API
├── utils/           
│   ├── auth.go            # Login & token storage
│   ├── pdf.go             # Generates PDF reports from student data
```


## 🔍 Troubleshooting

Error: tokens not found in login response cookies

Possible causes:

Backend sends cookies with Secure flag over HTTP

Backend not running

Fix:

Remove Secure flag from backend cookies for local dev OR run backend over HTTPS.

Double-check .env values.

Check cookies manually:

Test login API in Postman and inspect Set-Cookie headers.

🌐 Endpoints

GET ```/ping```

Description: Quick health check to confirm the service is running.

Response:
```pong```

GET ```/api/v1/students/{id}/report```

Description: Generates a PDF report for the student with the given ID.

URL Parameters:

id – Student ID in the backend database.

Response Headers:

```bash
Content-Type: application/pdf
Content-Disposition: attachment; filename="student_report_<id>.pdf"
```

Example:

```bash
 curl -X GET http://localhost:8080/api/v1/students/1/report --output student_report_1.pdf
```

For backend documentation, see [../backend/README.md](../backend/README.md)

For frontend documentation, see [../frontend/README.md](../frontend/README.md)