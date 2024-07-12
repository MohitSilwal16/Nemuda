package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/MohitSilwal16/Nemuda/server/models"
	"github.com/MohitSilwal16/Nemuda/server/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var sqlDB *sql.DB

func Init_MariaDB() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	dbUser := os.Getenv("sqlDBUser")
	dbPass := os.Getenv("sqlDBPass")
	dbName := os.Getenv("sqlDBName")
	dbPort := os.Getenv("sqlDBPort")

	if dbUser == "" || dbName == "" || dbPass == "" {
		return errors.New("DATABASE NAME, USER NOT SPECIFIED IN .ENV FILE")
	}

	// On port 3306 MYSQL is running
	// username:password@tcp(localhost:3306)/mydb

	dbURL := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s", dbUser, dbPass, dbPort, dbName)

	sqlDB, err = sql.Open("mysql", dbURL)

	if err != nil {
		log.Println("Error:", err)
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
func IsSessionTokenValid(token string) (bool, error) {
	rows, err := sqlDB.Query("SELECT 1 from users WHERE Token = ?;", token)

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
		isSessionTokenDuplicate, err := IsSessionTokenValid(sessionToken)
		if err != nil {
			return "", err
		}
		if !isSessionTokenDuplicate {
			break
		}
		sessionToken = utils.TokenGenerator()
		log.Println("Duplicate Token " + sessionToken)
	}

	_, err := sqlDB.Query("UPDATE users SET Token = ? WHERE Username = ?;", sessionToken, username)

	if err != nil {
		return "", nil
	}
	return sessionToken, nil
}

func DeleteTokenInDB(sessionToken string) error {
	isSessionTokenValid, err := IsSessionTokenValid(sessionToken)
	if err != nil {
		return err
	}

	// Return error if session token is not present in db
	if !isSessionTokenValid {
		return errors.New("INVALID SESSION TOKEN")
	}

	stmt, err := sqlDB.Prepare("UPDATE users SET TOKEN = '' WHERE Token = ?;")

	if err != nil {
		return err
	}
	_, err = stmt.Exec(sessionToken)

	if err != nil {
		return err
	}
	return nil
}

func DoesUserExists(username string) (bool, error) {
	rows, err := sqlDB.Query("SELECT 1 FROM users WHERE Username = ?;", username)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	// Returns true if it has any data else false
	return rows.Next(), nil
}

func AddUser(user models.User) error {
	userAlreadyExists, err := DoesUserExists(user.Username)

	if err != nil {
		return err
	}

	if userAlreadyExists {
		return errors.New("USERNAME IS ALREADY USED")
	}

	stmt, err := sqlDB.Prepare("INSERT INTO users (Username , Password , Token) VALUE (? , ? , ?);")

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
	doesUserExists, err := DoesUserExists(user.Username)

	if err != nil {
		return false, err
	}

	if !doesUserExists {
		return false, errors.New("USER DOESN'T EXISTS")
	}

	rows, err := sqlDB.Query("SELECT 1 FROM users WHERE Username = ? AND Password = ?;", user.Username, user.Password)

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

	rows, err := sqlDB.Query("SELECT Username FROM users WHERE Token = ? LIMIT 1;", sessionToken)

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

func GetMessages(sender string, receiver string) ([]models.Message, error) {
	rows, err := sqlDB.Query("SELECT * FROM messages WHERE (Sender = ? AND Receiver = ? ) OR (Sender = ? AND Receiver = ? ) ORDER BY DateTime;", sender, receiver, receiver, sender)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message

		if err = rows.Scan(&message.Sender, &message.Receiver, &message.MessageContent, &message.Status, &message.DateTime); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func AddMessage(message models.Message) error {
	stmt, err := sqlDB.Prepare("INSERT INTO messages (Sender, Receiver, MessageContent, Status, DateTime) VALUE (? , ? , ?, ?, ?);")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(message.Sender, message.Receiver, message.MessageContent, message.Status, message.DateTime)

	if err != nil {
		return err
	}
	return nil
}

// Return users whose name starts with searchString
func SearchUsersByPattern(searchString string) ([]string, error) {
	searchString += "%"

	rows, err := sqlDB.Query("SELECT Username FROM users WHERE Username LIKE ?;", searchString)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var usernames []string
	for rows.Next() {
		var username string

		if err = rows.Scan(&username); err != nil {
			return nil, err
		}
		usernames = append(usernames, username)
	}

	return usernames, nil
}

func ChangeStatusOfMessage(message models.Message, newStatus string) error {
	result, err := sqlDB.Exec("UPDATE messages SET Status = ? WHERE Sender = ? AND Receiver = ? AND MessageContent = ? AND DateTime = ?;", newStatus, message.Sender, message.Receiver, message.MessageContent, message.DateTime)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("MESSAGE NOT FOUND")
	}

	return nil
}
