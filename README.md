# 🚀 Feature Flag Service

A lightweight and efficient **Feature Flag Service** built with **Golang, Gin, PostgreSQL, Redis, and Swagger API Documentation**. This service allows teams to manage and toggle feature flags dynamically for better **continuous deployment** and **A/B testing** strategies.

## 🌟 Features
- ✅ **JWT Authentication** (User Registration & Login)
- ✅ **Feature Flag Management** (Create, Read, Update, Delete)
- ✅ **PostgreSQL Database Integration**
- ✅ **Redis for Caching**
- ✅ **Swagger API Documentation**
- ✅ **Dockerized for Easy Deployment**
- ✅ **Mocked Database for CI/CD Testing**

---

## 📂 Project Structure
```
feature-flag-service/
│-- internal/
│   │-- config/          # Database & Redis Configuration
│   │-- handlers/        # API Route Handlers
│   │-- middleware/      # Authentication Middleware
│   │-- models/         # Database Models
│   │-- tests/          # Unit & Integration Tests
│-- docs/               # Swagger Documentation
│-- main.go             # Entry Point
│-- Dockerfile          # Docker Build Config
│-- docker-compose.yml  # Docker Compose Services
│-- README.md           # Project Documentation
```

---

## 🔧 Setup & Installation
### **1️⃣ Clone the Repository**
```sh
git clone https://github.com/your-username/feature-flag-service.git
cd feature-flag-service
```

### **2️⃣ Set Up Environment Variables**
Create a **.env** file in the root directory:
```sh
PORT=8080
DATABASE_URL=postgres://postgres:password@postgres:5432/feature_flags?sslmode=disable
REDIS_URL=redis:6379
JWT_SECRET=your_secret_key
```

### **3️⃣ Run the Application Locally**
#### **With Docker (Recommended)**
```sh
docker-compose up --build
```

#### **Without Docker**
Ensure PostgreSQL and Redis are running, then:
```sh
go mod tidy
go run main.go
```

---

## 📌 API Endpoints
### **🔑 Authentication**
| Method | Endpoint       | Description           |
|--------|--------------|----------------------|
| POST   | `/register`  | Register a new user  |
| POST   | `/login`     | Authenticate & get JWT |

### **🚀 Feature Flags**
| Method | Endpoint           | Description                     |
|--------|------------------|--------------------------------|
| POST   | `/api/flags`      | Create a new feature flag      |
| GET    | `/api/flags`      | Get all feature flags          |
| GET    | `/api/flags/{id}` | Get a single feature flag by ID |
| PUT    | `/api/flags/{id}` | Update a feature flag          |
| DELETE | `/api/flags/{id}` | Delete a feature flag          |

**📖 Swagger Documentation**
- Once the service is running, access Swagger UI:
  👉 [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## ✅ Running Tests
To run tests **with database mocking**:
```sh
TEST_MODE=true go test ./internal/tests -v
```

For **full integration tests**:
```sh
go test ./...
```

---

## 📦 Deployment
### **🚀 Deploy with Railway**
1. Install Railway CLI:
   ```sh
   curl -fsSL https://railway.app/install.sh | sh
   ```
2. Initialize project:
   ```sh
   railway init
   ```
3. Deploy:
   ```sh
   railway up
   ```

---

## 🤝 Contributing
We welcome contributions! Feel free to **open an issue** or **submit a pull request**.

---

## 🛡️ License
This project is **MIT Licensed**. See `LICENSE` for details.

---

## 🌟 Acknowledgments
- **Golang & Gin** for backend development
- **Swagger** for API Documentation
- **PostgreSQL & Redis** for database management
- **Railway & Docker** for seamless deployment