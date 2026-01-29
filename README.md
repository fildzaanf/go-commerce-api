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

## ✨ Features

#### User Management

| Feature                   | Description                                                        |
| ------------------------- | ------------------------------------------------------------------ |
| User Registration & Login | Allows users to register and log in to access the platform         |
| Profile                   | Provides functionality to retrieve user profile information by ID  |

#### Product Management

| Feature           | Description                                                                    |
| ----------------- | ------------------------------------------------------------------------------ |
| Create Product    | Enables adding new products to the platform                                    |
| Update Product    | Allows updating existing product details by product ID                         |
| Delete Product    | Supports removing products from the platform by product ID                     |
| Retrieve Product  | Provides access to a single product by ID or a list of all available products  |

#### Payment Management

| Feature             | Description                                                             |
| ------------------- | ----------------------------------------------------------------------- |
| Create Payment      | Allows users to create new payments for products or services            |
| Retrieve Payment    | Provides access to all payments or details of a specific payment by ID  |
| Integration Payment | Supports real-time payment updates via Midtrans Webhook integration     |

## 📡 API Endpoints

#### Users

| Method | Endpoint        | Description                 |
| ------ | --------------- | --------------------------- |
| POST   | /users/register | Register a new user         |
| POST   | /users/login    | Login user                  |
| GET    | /users/:id      | Retrieve user profile by ID |

#### Products

| Method | Endpoint      | Description                    |
| ------ | ------------- | ------------------------------ |
| POST   | /products     | Create a new product           |
| PUT    | /products/:id | Update product details by ID   |
| DELETE | /products/:id | Delete product by ID           |
| GET    | /products/:id | Retrieve product details by ID |
| GET    | /products     | Retrieve all products          |

#### Payments

| Method | Endpoint                   | Description                                  |
| ------ | -------------------------- | -------------------------------------------- |
| POST   | /payments                  | Create a new payment                         |
| GET    | /payments/:id              | Retrieve payment details by ID               |
| GET    | /payments                  | Retrieve all payments                        |
| POST   | /payments/midtrans/webhook | Receive Midtrans webhook for payment updates |

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



