# 🚀 api.bruno-guimaraes.com

A robust, modern API built with Go for efficient email communications and extensible functionality.

---

## 🌟 Features

- ✉️ **Secure Email Processing**: Integrated with Resend for reliable email delivery
- 🔒 **Environment-Based Error Handling**: Context-aware error management
- ⚡ **High Performance**: Leveraging Go's concurrency and efficiency
- 🛡️ **CORS Protection**: Secure cross-origin resource sharing
- 🔐 **Token-Based Authentication**: Protected API endpoints
- 🎯 **Clean Architecture**: Modular and maintainable codebase

---

## 🛠️ Tech Stack

- **Language**: Go (v1.24+)
- **Containerization**: Docker
- **Orchestration**: Docker Compose
- **Email Service**: Resend
- **Environment Management**: `.env` configuration

---

## 🏗️ Project Architecture

Built following clean architecture principles with clear separation of concerns:

```
├── .github/         # GitHub CI/CD setup
├── application/     # API logic
├── domain/          # Business logic
├── errors/          # Error handling
├── infrastructure/  # External services
├── interfaces/      # API handlers and middleware
```

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (1.24 or higher)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### 🔧 Environment Variables

Create a `.env` file in the root directory with these variables:

```env
RESEND_API_KEY=your_resend_api_key   
AUTH_TOKEN=your_auth_token                
EMAIL_ADMIN=admin@example.com               
CORS_ORIGIN=https://your-allowed-origin.com 
GO_ENV=development
Environment (development/production)
```

### 🐳 Running with Docker

1. Build and start the containers:
```bash
docker-compose up -d
```

2. View container logs:
```bash
docker-compose logs -f
```

3. Stop the containers:
```bash
docker-compose down
```

### Running Locally

1. Install dependencies:
```bash
go mod tidy
```

2. Run the application:
```bash
go run main.go
```

The API will be available at `http://localhost:8080`.

---

## 🔒 Security Features

- **Environment-Based Error Handling**: Prevents sensitive data leaks
- **Secure Email Validation**: Sanitizes and validates email inputs
- **Protected Routes**: Authentication middleware for secure access
- **CORS Protection**: Restricts access to specified origins

---

## ⚙️ API Endpoints

### Send Email
- **Endpoint**: `POST /api/send-message`
- **Description**: Sends an email via the Resend service
- **Authentication**: Required (Bearer token)

#### Request Body
```json
{
  "senderName": "John Doe",
  "senderEmail": "john@example.com",
  "content": "Hello, this is a test message"
}
```

#### Response
- **Success (200)**:
```json
{
    "adminEmail":"token number",
    "clientEmail":"token number"
    }

```
- **Error (400/401/500)**:
```json
{
  "error": "Error description"
}
```

---

## 🧪 Testing

Run the test suite:
```bash
go test ./... -v
```

---

## 📬 Contact

- **Email**: [bruno.sil16441@gmail.com](mailto:bruno.sil16441@gmail.com)
- **Website**: [bruno-guimaraes.com](https://bruno-guimaraes.com)
- **LinkedIn**: [linkedin.com/in/bruno-webdev](https://linkedin.com/in/bruno-webdev)
- **GitHub**: [github.com/BrunoGuimaraesSilva](https://github.com/BrunoGuimaraesSilva)


