
# Personal Finance Management Backend

A backend system for personal finance management built with **Go (Golang)** and the **Gin Framework**, designed to provide a fast and flexible API. 


## Features

- User registration and login with JWT Authentication
- Manage categories (Categories)
- Set and track budgets (Budgets)
- Record transactions (Transactions)
- Generate monthly financial reports (Monthly Reports)
- Middleware for user authorization
- MongoDB integration for data storage

## Tech Stack

- **Language**: Go (Golang)
- **Framework**: Gin
- **Database**: MongoDB
- **Authentication**: JSON Web Token (JWT)


## Installation

### 1. Clone this repository

```bash
  git clone https://github.com/DomePannawat/Personal-Finance-Management-Backend.git
```
```bash
  cd Personal-Finance-Management-Backend
```
    
### 2. Configure environment variables
```bash
MONGODB_URI=mongodb+srv://<your-mongodb-uri>
MONGODB_DB=personal_finance
PORT=<Any PORT number can be used>

```

### 3. Install dependencies
```bash
  go mod tidy

```

### 4. Run the project
```bash
  go run main.go

```
