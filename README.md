# 💰 Wallet API

A simple wallet management API built using **Go (Gin framework)** and **PostgreSQL**, providing basic functionality for withdrawing funds from a user’s account.  
The project includes **database migrations with seed data**, making it easy to start testing right after setup.

---

## 📋 Table of Contents
1. [Assumptions](#assumptions)
2. [Process Flow](#process-flow)
3. [Tech Stack](#tech-stack)
4. [Prerequisites](#prerequisites)
5. [Database Architecture](#database-architecture)
6. [Setup & Run](#setup--run)
7. [API Contract](#api-contract)
8. [Limitations & Future Improvements](#limitations--future-improvements)

---

## 📌 Assumptions

- Users already exist in the system (seeded during migration).
- Withdrawals are only allowed if the user has sufficient balance.
- Transactions are immediately marked as completed for simplicity.
- No user registration/login flow is implemented for now — focus is on withdrawal logic.

---

## 📌 Process Flow

1. **Clone the Repository**
   - Download the project from GitHub to your local machine.

2. **Install Dependencies**
   - Vendor dependencies with `go mod vendor` and clean up with `go mod tidy`.

3. **Setup Environment**
   - Copy `.env.example` to `.env` and adjust variables according to your setup (local PostgreSQL or Docker PostgreSQL).

4. **Prepare Database**
   - Create the database in PostgreSQL (`wallet_db` by default).
   - The app uses **golang-migrate** for migrations and **sqlc** to generate SQL queries.
   - Migrations run automatically when the app starts.
   - **Seed data** for `user` table is included, so you can directly perform withdrawals.

5. **Run the Application**
   - Start the app with:
     ```bash
     go run cmd/main.go
     ```

6. **Test APIs**
   - Use **Postman** or any API client to test endpoints.
   - Example: `POST /withdraw` with an existing seeded user ID.

---

## 🛠 Tech Stack

- **Go** v1.24.6
- **Gin** (HTTP framework)
- **PostgreSQL** v16
- **golang-migrate** (database migrations)
- **sqlc** (SQL code generation)
- **DBeaver** (GUI tool for PostgreSQL)
- **Docker** (optional, for running PostgreSQL)
- **Postman** (API testing)
- **VS Code** (editor)

---

## 📦 Prerequisites

Before running the application, ensure you have:

- [Go](https://go.dev/) >= 1.24.6
- [PostgreSQL](https://www.postgresql.org/) >= v16  
  (Can run locally or via Docker)
- [golang-migrate CLI](https://github.com/golang-migrate/migrate)
- [sqlc CLI](https://sqlc.dev/)
- [DBeaver](https://dbeaver.io/) (for database management)
- Git
- (Optional) Docker & Docker Compose

---

## 🗄 Database Architecture Diagram

The following Entity-Relationship Diagram (ERD) illustrates the database structure for the Wallet API:
![Wallet API ERD](docs/wallet_api_erd.png)

---

## 🚀 Setup & Run

1. **Clone the repository**
   ```bash
   git clone https://github.com/vincentweilasto16/wallet-api.git
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   go mod vendor
   ```
   
3. Setup environment
   ```bash
   cp .env.example .env
   ```
   Adjust *.env* to match your local PostgreSQL or Docker Setup

4. Create the database
   ```bash
   CREATE DATABASE wallet_db;
   ```

5. Run the application
   ```bash
   go run cmd/main.go
   ```

6. Test the API with Postman

---

## 📜 API Contract

1. **Get User By ID**
   Endpoint:
   ```bash
   GET /api/v1/users/:id
   ```

   Response (Success - 200 OK):
   ```bash
   {
    "data": {
        "id": "67e1e382-7122-4e77-b47f-f4e940cbf385",
        "name": "Andi Wijaya",
        "email": "andi.wijaya@example.com",
        "balance": 1500000,
        "created_at": "2025-08-15T17:17:21.515714+07:00",
        "updated_at": "2025-08-15T17:17:21.515714+07:00",
        "deleted_at": null
    },
    "meta": {
        "status": 200,
        "message": "ok"
    }
   }
   ```

   Response (Failure - User Not Found - 404 Not Found):
   ```bash
    {
        "data": null,
        "errors": [
            {
                "code": "NOT_FOUND",
                "message": "user not found"
            }
        ],
        "meta": {
            "status": 404,
            "message": "user not found"
        }
    }
   ```

2. **Withdraw**
   Endpoint:
   ```bash
   POST /api/v1/transactions/withdraw
   ```

   Request Body:
   ```bash
   {
    "user_id": "a440fd30-6894-4e93-b041-2f577c09d002",
    "amount": 20000
   }
   ```

   Response (Success - 200 OK):
   ```bash
   {
        "data": {
            "message": "Withdrawal successful"
        },
        "meta": {
            "status": 200,
            "message": "OK"
        }
   }
   ```

   Response (Failure - Insufficient Balance - 422 Unprocessable entity):
   ```bash
    {
        "data": null,
        "errors": [
            {
                "code": "UNPROCESSABLE_ENTITY",
                "message": "insuficient balance"
            }
        ],
        "meta": {
            "status": 422,
            "message": "insuficient balance"
        }
    }
   ```

   Response (Failure - User Not Found - 404 Not Found):
   ```bash
    {
        "data": null,
        "errors": [
            {
                "code": "NOT_FOUND",
                "message": "user not found"
            }
        ],
        "meta": {
            "status": 404,
            "message": "user not found"
        }
    }
   ```

   

---

## 🔮 Limitations & Future Improvements

- **Authentication & Authorization**
  - Implement JWT with Bearer tokens for all API endpoints.

- **Withdraw Transaction Handling**
  - Add rollback mechanism for failed withdrawals.
  - Implement idempotency to prevent duplicate withdrawals.
  - Add rate limiting for withdrawals.

- **Environments**
  - Separate configuration for development, staging, and production.

- **Testing & CI/CD*
  - Unit tests for controller and service layers.
  - GitHub Actions workflow for build & test.

- **Deployment**
  - Dockerfile for containerization.
  - Kubernetes manifests for deployment.

- **Monitoring**
  - Currently only basic logging; can be improved with Prometheus & Grafana.
 
- **Error Handling**
  - Could be enhanced with better structured error responses.

- **Feature Enhancements**
  - Expand transaction status (pending, refund, canceled, failed).
  - Integrate third-party payment gateways.
  - Prevent withdrawals when balance is zero (backend & frontend validation).

---

## 📦 Note

When the application runs migrations, it automatically:
1. Creates required tables.
2. Inserts seed data into the user table.
   
This means you can call the withdraw API immediately with a seeded user_id.

---


