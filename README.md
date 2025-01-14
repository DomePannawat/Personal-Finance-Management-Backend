
# Personal Finance Management Backend

Developed a system to help individuals manage their personal finances effectively. 
This includes features like user registration, budget tracking, recording expenses, 
and generating financial reports. The system was built using Go (Golang) and the Gin Framework.


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

## License

[![Apache-2.0 License](https://img.shields.io/badge/license-Apache%202.0-blue?style=flat-square)](https://www.apache.org/licenses/LICENSE-2.0)
