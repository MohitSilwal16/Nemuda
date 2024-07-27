## To get started with this project, you'll need to set up environment variables in two files: `main.env` and `.env`.

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

