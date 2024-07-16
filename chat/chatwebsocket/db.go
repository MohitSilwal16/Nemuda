package chatwebsocket

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

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
		return errors.New("DATABASE NAME, USER & PASS NOT SPECIFIED IN .ENV FILE")
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

func isSessionTokenValid(client *Client) {
	if client == nil {
		log.Println("Null client\nSource: isSessionTokenValid()")
		return
	}

	if client.SessionToken == "" {
		client.SendMessage <- Message{
			Error: "Invalid Session Token",
		}
		return
	}

	rows, err := sqlDB.Query("SELECT 1 from users WHERE Token = ?;", client.SessionToken)

	if err != nil {
		log.Println("Error:", err, "\nSource: isSessionTokenValid()")
		log.Println("Description: Cannot validate session token")

		client.SendMessage <- Message{
			Error: "Internal Server Error",
		}
		return
	}

	defer rows.Close()

	// Return false cuz Sesion Token is Invalid
	if !rows.Next() {
		client.SendMessage <- Message{
			Error: "Invalid Session Token",
		}
	}
}

func addMessageInDB(client *Client, message Message) {
	if client == nil {
		log.Println("Null client\nSource: addMessage()")
		return
	}

	tableName := "messages_" + message.Receiver
	// Tablename can't be used as placeholder
	stmt, err := sqlDB.Prepare("INSERT INTO " + tableName + " (Sender, Receiver, MessageContent, Status, DateTime) VALUE (? , ? , ?, ?, ?);")

	if err != nil {
		log.Println("Error:", err, "\nSource: addMessage()")
		log.Println("Description: Cannot prepare statement")

		client.SendMessage <- Message{
			Error: "Internal Server Error",
		}
		return
	}

	_, err = stmt.Exec(message.Sender, message.Receiver, message.MessageContent, message.Status, message.DateTime)

	if err != nil {
		log.Println("Error:", err, "\nSource: addMessage()")
		log.Println("Description: Cannot execute statement")

		client.SendMessage <- Message{
			Error: "Internal Server Error",
		}
		return
	}
}

func changeStatusOfMessage(client *Client, sender string, receiver string, newStatus string) {
	if client == nil {
		log.Println("Null client\nSource: changeStatusOfMessage()")
		return
	}

	var err error
	if newStatus == "Read" {
		tableName := "messages_" + sender
		_, err = sqlDB.Exec("UPDATE "+tableName+" SET Status = 'Read' WHERE Sender = ?;", receiver)
	} else if newStatus == "Delivered" {
		tableName := "messages_" + sender
		_, err = sqlDB.Exec("UPDATE " + tableName + " SET Status = 'Delivered' WHERE Status = 'Sent';")
	}

	if err != nil {
		log.Println("Error:", err, "\nSource: changeStatusOfMessage()")
		log.Println("Description: Cannot change status of message")

		client.SendMessage <- Message{
			Error: "Internal Server Error",
		}
		return
	}
}

func getUsernameBySessionToken(sessionToken string) (string, error) {
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
