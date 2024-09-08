# Blogging Platform Integrated with Real-Time Messaging 

## Project's Website

The Project's Webpage might be live at [http://nemuda.hopto.org](http://nemuda.hopto.org). Please check it out for more details.
<br>
### NOTE THAT THE SITE MAY BE TEMPORARILY DOWN.

---

## Overview

This is a Blogging platform integrated with Real-Time Messaging, combining multiple components to deliver an adaptable user experience.

- **Frontend Web Server**: Written in Go using Gin Gonic, handles HTTP requests from Users through Browser & communicates with the Backend Server via gRPC.
- **Backend Server**: Written in Go, handles logic & data processing, communicates with the Frontend Server & Flutter Application using gRPC.
- **Messaging Server**: Written in Go, manages Real-Time Messaging using Channels & Goroutines, communicates with the Frontend Webpage & Flutter Application using WebSocket.
- **Flutter Application**: Manages state using BLoC pattern & interacts with the Backend Server via gRPC & Messaging Server via WebSocket Connection.
- **Data Storage**: MySQL, MongoDB, & S3 Bucket (for images & apk file).
- **Hosting**: AWS EC2.

## Features

- **Blog Management**: Create, Read, Update, Delete, Like, Comment Blog Posts, Filter Blogs by Tag & On Demand Loading of Blogs.
- **Real-Time Messaging**: Handle Real-Time Messaging & On Demand Loading of Previous Messages with Read Receipts.
<br>

---

# To get started with this project, you'll need to set up environment variables in: `main.env` & `.env`.

## 1. Create `main.env`

Create a file named `main.env` in the server directory:
```dotenv
# MySQL Configuration for Authentication & Fetching Previous Messages

# MySQL Database Name
sqlDBName="your-database-name"

# MySQL User
sqlDBUser="your-database-username"

# MySQL Password
sqlDBPass="your-database-password"

# MySQL Port (default is 3306)
sqlDBPort="3306"

# MongoDB Configuration to Store Blogs

# MongoDB Database Name
mongoDBName="your-mongodb-database-name"

# MongoDB Collection Name
mongoCollectionName="your-mongodb-collection-name"
```

## 2. Create `.env`

Create a file named `.env` in the server directory:
```dotenv
# AWS Configuration for Connecting the Back-End Server to an S3 Bucket for Storing Blog Images

AWS_REGION="your-aws-region"                   # e.g., us-east-1
AWS_ACCESS_KEY_ID="your-access-key-id"         # e.g., AKIAIOSFODNN7EXAMPLE
AWS_SECRET_ACCESS_KEY="your-secret-access-key" # e.g., wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
```

## 3. Create `.env` again

Create a file named `.env` again in the chat directory:
```dotenv
# Database Configuration for Authentication & Performing Operations on Messages (Add Message, Change Status of Message)

# MySQL Database Name
sqlDBName="your-database-name"

# MySQL User
sqlDBUser="your-database-username"

# MySQL Password
sqlDBPass="your-database-password"

# MySQL Port (default is 3306)
sqlDBPort="3306"
```

# To build and run the project(In Linux):

```bash
make build # Run this command to compile the code.
make run
```

# To stop this Project(In Linux):

```bash
make kill
```
