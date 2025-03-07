# go-commerce-api

## 📌 Overview
This project is a simple e-commerce system that includes key features such as user management (sellers and buyers), product catalog, and payment system using Midtrans payment gateway.

## 🚀 Tools and Technologies 
- Go Programming Language
- Echo Framework
- GORM
- MySQL / PostgreSQL
- Docker
- JWT Authentication
- Midtrans Payment Gateway
- Amazon Simple Storage Service (S3)
- Simple Mail Transfer Protocol (SMTP)

## 📂 Folder Structure
```
/go-commerce-api
├── cmd/                # application entry point
├── docs/               # application documentation
├── infrastructure/     # infrastructure configurations such as database
├── internal/           # core business logic of the application
│   ├── payment/        # payment module
│   ├── product/        # product module
│   ├── user/           # user module
│   │   ├── domain/     # domain definition or business model
│   │   ├── dto/        # data transfer object for request & response
│   │   ├── entity/     # entity for database model representation
│   │   ├── handler/    # layer for handling HTTP requests
│   │   ├── repository/ # layer for database access
│   │   ├── router/     # routing configuration
│   │   ├── service/    # layer for business logic
├── pkg/                # libraries or reusable helper functions
├── .env.example        # example configuration file for environment variables
├── Dockerfile          # docker configuration file for containerization
├── .gitignore          # list of files ignored by git
├── go.mod              # Go module dependencies
├── go.sum              # Go dependencies checksum
├── README.md           # project documentation
```

## 🛠️ Installation & Running the Project
### 1️⃣ Prerequisites
Make sure you have installed:
- [Go](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/) / [MySQL](https://dev.mysql.com/downloads/)

### 2️⃣ Clone the Repository
```bash
git clone <repo-url>
cd <repo-directory>
```

### 3️⃣ Configure Environment
Create a `.env` file based on `.env.example` and place it in the `cmd` directory.

### 4️⃣ Run the Application
```bash
cd cmd
go run main.go
```

## 🎯 Git Flow & Branching Strategy
1. `main` → the main branch used only for production releases
2. `development` → the primary branch for feature development before merging into main
3. `feature/xyz` → branch for new feature development (replace `xyz` with the feature name)
4. `bugfix/xyz` → branch for bug fixes (replace `xyz` with the bug name)