# Blogging Platform Integrated with Real-Time Messaging 

## Overview

This project is a blogging platform with integrated real-time messaging, combining multiple components to deliver an interactive and adaptable user experience.

- **Frontend Server**: Written in Go using Gin Gonic, handles HTTP requests from users and communicates with the backend server via gRPC.
- **Backend Server**: Written in Go, handles business logic and data processing, communicates with the frontend server & flutter application using gRPC.
- **Messaging Server**: Written in Go, manages real-time messaging and concurrency using channels and goroutines, communicates with the frontend server & flutter application using gRPC.
- **Flutter Application**: Manages state using BLoC pattern and interacts with the backend and messaging servers via gRPC.
- **Data Storage**: MySQL, MongoDB, and S3 Bucket (for images).
- **Hosting**: AWS EC2.

## Features

- **Blog Management**: Create, read, update, delete, like, comment blog posts.
- **Real-Time Messaging**: Handle real-time messaging with efficient concurrency.
<br>

---

# To get started with this project, you'll need to set up environment variables in: `main.env` and `.env`.

## 1. Create `main.env`

Create a file named `main.env` in the server directory:
```dotenv
# Database Configuration

# MySQL Database Name
sqlDBName="your-database-name"

# MySQL User
sqlDBUser="your-database-username"

# MySQL Password
sqlDBPass="your-database-password"

# MySQL Port (default is 3306)
sqlDBPort="3306"

# MongoDB Configuration

# MongoDB Database Name
mongoDBName="your-mongodb-database-name"

# MongoDB Collection Name
mongoCollectionName="your-mongodb-collection-name"
```

## 2. Create `.env`

Create a file named main.env in the server directory:
```dotenv
AWS_REGION="your-aws-region"           # e.g., us-east-1
AWS_ACCESS_KEY_ID="your-access-key-id" # e.g., AKIAIOSFODNN7EXAMPLE
AWS_SECRET_ACCESS_KEY="your-secret-access-key" # e.g., wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
```

## 3. Create `.env` again

Create a file named `.env` again in the chat directory:
```dotenv
# Database Configuration

# MySQL Database Name
sqlDBName="your-database-name"

# MySQL User
sqlDBUser="your-database-username"

# MySQL Password
sqlDBPass="your-database-password"

# MySQL Port (default is 3306)
sqlDBPort="3306"
```

# To run Back-end server:
- cd server
- go mod tidy
- go run main.go

# To run Front-end Server:
- cd client
- go mod tidy
- go run main.go

# To run Messaging Server:
- cd chat
- go mod tidy
- go run main.go
