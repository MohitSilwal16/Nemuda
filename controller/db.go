package controller

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/MohitSilwal16/Nemuda/models"
	"github.com/MohitSilwal16/Nemuda/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var sqlDB *sql.DB

func init() {
	utils.ClearScreen()

	err := godotenv.Load("dotenv")

	if err != nil {
		panic(err)
	}

	dbUser := os.Getenv("dbUser")
	dbPass := os.Getenv("dbPass")
	dbName := os.Getenv("dbName")

	if dbUser == "" || dbName == "" || dbPass == "" {
		panic("Provide database user , pass & database name")
	}

	// On port 3306 MYSQL is running
	// username:password@tcp(localhost:3306)/mydb

	dbURL := fmt.Sprintf("%s:%s@tcp(localhost:3305)/%s", dbUser, dbPass, dbName)

	sqlDB, err = sql.Open("mysql", dbURL)

	if err != nil {
		panic(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		fmt.Println("Connection isn't established")
		panic(err)
	}
	fmt.Println("Connection with database is established")
}

// Helper methods
func checkDuplicateToken(token string) bool {
	rows, err := sqlDB.Query("SELECT 1 from Users WHERE Token = ?;", token)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	// Returns true if it has any data else false
	return rows.Next()
}

func SearchUserByName(username string) bool {
	rows, err := sqlDB.Query("SELECT 1 FROM Users WHERE Username = ?;", username)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	// Returns true if it has any data else false
	return rows.Next()
}

func AddUser(user models.User) {
	sessionToken := utils.TokenGenerator()

	for checkDuplicateToken(sessionToken) {
		sessionToken = utils.TokenGenerator()
		fmt.Println("Duplicate Token")
	}

	stmt, err := sqlDB.Prepare("INSERT INTO Users (Username , Password , Token) VALUE (? , ? , ?);")

	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(user.Username, user.Password, sessionToken)

	if err != nil {
		panic(err)
	}
}

func VerifyIdPass(user models.User) bool {
	rows, err := sqlDB.Query("SELECT 1 FROM Users WHERE Username = ? AND Password = ?;", user.Username, user.Password)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	return rows.Next()
}
