package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/MohitSilwal16/Nemuda/models"
	"github.com/MohitSilwal16/Nemuda/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var PORT = "6969"

var sqlDB *sql.DB

func Init_MariaDB() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	dbUser := os.Getenv("sqlDBUser")
	dbPass := os.Getenv("sqlDBPass")
	dbName := os.Getenv("sqlDBName")

	if dbUser == "" || dbName == "" || dbPass == "" {
		return errors.New("DATABASE NAME, USER & PASSWORD NOT SPECIFIED IN .ENV FILE")
	}

	// On port 3306 MYSQL is running
	// username:password@tcp(localhost:3306)/mydb

	dbURL := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s", dbUser, dbPass, PORT, dbName)

	sqlDB, err = sql.Open("mysql", dbURL)

	if err != nil {
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Println("Connection with Maria DB isn't established")
		return err
	}
	log.Println("Connection with Maria DB is established")
	return nil
}

// Helper methods
func IsSessionTokenDuplicate(token string) (bool, error) {
	rows, err := sqlDB.Query("SELECT 1 from Users WHERE Token = ?;", token)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	// Returns true if it has any data else false
	return rows.Next(), nil
}

func UpdateTokenInDBAndReturn(username string) (string, error) {
	// Generate session token
	sessionToken := utils.TokenGenerator()

	// Checking whether the generated session token is already used or not
	// If so then generate another session token & keep this loop until we find a unique session token
	for {
		isSessionTokenDuplicate, err := IsSessionTokenDuplicate(sessionToken)
		if err != nil {
			return "", err
		}
		if !isSessionTokenDuplicate {
			break
		}
		sessionToken = utils.TokenGenerator()
		log.Println("Duplicate Token " + sessionToken)
	}

	_, err := sqlDB.Query("UPDATE Users SET Token = ? WHERE Username = ?;", sessionToken, username)

	if err != nil {
		return "", nil
	}
	return sessionToken, nil
}

func DeleteTokenInDB(sessionToken string) error {
	isSessionTokenValid, err := IsSessionTokenDuplicate(sessionToken)
	if err != nil {
		return err
	}

	// Return error if session token is not present in db
	if !isSessionTokenValid {
		return errors.New("INVALID SESSION TOKEN")
	}

	stmt, err := sqlDB.Prepare("UPDATE Users SET TOKEN = '' WHERE Token = ?;")

	if err != nil {
		return err
	}
	_, err = stmt.Exec(sessionToken)

	if err != nil {
		return err
	}
	return nil
}

func SearchUserByName(username string) (bool, error) {
	rows, err := sqlDB.Query("SELECT 1 FROM Users WHERE Username = ?;", username)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	// Returns true if it has any data else false
	return rows.Next(), nil
}

func AddUser(user models.User) error {
	userAlreadyExists, err := SearchUserByName(user.Username)

	if err != nil {
		return err
	}

	if userAlreadyExists {
		return errors.New("USERNAME IS ALREADY USED")
	}

	stmt, err := sqlDB.Prepare("INSERT INTO Users (Username , Password , Token) VALUE (? , ? , ?);")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Username, user.Password, user.Token)

	if err != nil {
		return err
	}
	return nil
}

func VerifyIdPass(user models.User) (bool, error) {
	doesUserExists, err := SearchUserByName(user.Username)

	if err != nil {
		return false, err
	}

	if !doesUserExists {
		return false, errors.New("USER DOESN'T EXISTS")
	}

	rows, err := sqlDB.Query("SELECT 1 FROM Users WHERE Username = ? AND Password = ?;", user.Username, user.Password)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	return rows.Next(), nil
}

func GetUsernameBySessionToken(sessionToken string) (string, error) {
	if sessionToken == "" {
		return "", errors.New("INVALID SESSION TOKEN")
	}

	rows, err := sqlDB.Query("SELECT Username FROM Users WHERE Token = ? LIMIT 1;", sessionToken)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	var username string
	if rows.Next() {
		err = rows.Scan(&username)

		if err != nil {
			return "", err
		}
		return username, nil
	} else {
		return "", errors.New("INVALID SESSION TOKEN")
	}
}
