package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/MohitSilwal16/Nemuda/server/pb"
	"github.com/MohitSilwal16/Nemuda/server/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var sqlDB *sql.DB

// Prepared Statements
var stmtIsSessionTokenValid *sql.Stmt
var stmtUpdateSessionToken *sql.Stmt
var stmtDeleteSessionToken *sql.Stmt // Set Session Token as ""
var stmtDoesUserExists *sql.Stmt
var stmtAddUser *sql.Stmt
var stmtVerifyIdPass *sql.Stmt
var stmtGetUserBySessionToken *sql.Stmt
var stmtSearchUsersByPattern *sql.Stmt

func Init_MariaDB() error {
	err := godotenv.Load("main.env")

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
	log.Println("Formulating Prepared Statements ...")

	stmtIsSessionTokenValid, err = sqlDB.Prepare("SELECT 1 FROM users WHERE Token = ?;")
	if err != nil {
		log.Println("Error: Failed to Create Prepared Statement for Validating Session Token")
		return err
	}

	stmtUpdateSessionToken, err = sqlDB.Prepare("UPDATE users SET Token = ? WHERE Username = ?;")
	if err != nil {
		log.Println("Error: Failed to Prepare Statement for Updating Session Token")
		return err
	}

	stmtDeleteSessionToken, err = sqlDB.Prepare("UPDATE users SET Token = '' WHERE Token = ?;")
	if err != nil {
		log.Println("Error: Failed to Prepare Statement for Deleting Session Token")
		return err
	}

	stmtDoesUserExists, err = sqlDB.Prepare("SELECT 1 FROM users WHERE Username = ?;")
	if err != nil {
		log.Println("Error: Failed to Prepare Statement for checking whether Username is used or not")
		return err
	}

	stmtAddUser, err = sqlDB.Prepare("INSERT INTO users (Username , Password , Token) VALUE (? , ? , ?);")
	if err != nil {
		log.Println("Error: Failed to Prepare Statement for Registering User")
		return err
	}

	stmtVerifyIdPass, err = sqlDB.Prepare("SELECT 1 FROM users WHERE Username = ? AND Password = ?;")
	if err != nil {
		log.Println("Error: Failed to Prepare Statement for Validating Username & Password")
		return err
	}

	stmtGetUserBySessionToken, err = sqlDB.Prepare("SELECT Username FROM users WHERE Token = ?;")
	if err != nil {
		log.Println("Error: Failed to Prepare Statement for getting Username from Session Token")
		return err
	}

	stmtSearchUsersByPattern, err = sqlDB.Prepare("SELECT Username FROM users WHERE Username LIKE ? LIMIT 10;")
	if err != nil {
		log.Println("Error: Failed to Prepare Statement for Searching Username by Pattern")
		return err
	}

	log.Println("Prepared statements have been finalized")
	return nil
}

func IsSessionTokenValid(token string) (bool, error) {
	// If a user logs outs then his cookie's value is set as empty string ""
	// So "" is always Invalid Token
	if token == "" {
		return false, nil
	}

	rows, err := stmtIsSessionTokenValid.Query(token)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	// Returns true if it has any data else false
	return rows.Next(), nil
}

func UpdateTokenInDBAndReturn(username string) (string, error) {
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

	_, err := stmtUpdateSessionToken.Exec(sessionToken, username)

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

	_, err = stmtDeleteSessionToken.Exec(sessionToken)
	if err != nil {
		return err
	}
	return nil
}

func DoesUserExists(username string) (bool, error) {
	rows, err := stmtDoesUserExists.Query(username)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	// Returns true if it has any data else false
	return rows.Next(), nil
}

func AddUser(user *pb.User) error {
	userAlreadyExists, err := DoesUserExists(user.Username)

	if err != nil {
		return err
	}

	if userAlreadyExists {
		return errors.New("USERNAME IS ALREADY USED")
	}

	_, err = stmtAddUser.Exec(user.Username, user.Password, user.Token)

	if err != nil {
		return err
	}

	tableName := "messages_" + strings.ToLower(user.Username)

	sqlQuery := fmt.Sprintf(`CREATE TABLE %s (
        Sender VARCHAR(20) NOT NULL CHECK (Sender <> ''),
        Receiver VARCHAR(20) NOT NULL CHECK (Receiver = '%s'),
        MessageContent VARCHAR(100) NOT NULL CHECK (MessageContent <> '' AND CHAR_LENGTH(MessageContent) <= 100),
        Status VARCHAR(10) NOT NULL CHECK (Status IN ('Sent', 'Delivered', 'Read')),
        DateTime DATETIME NOT NULL,
        FOREIGN KEY (Sender) REFERENCES users (Username),
        FOREIGN KEY (Receiver) REFERENCES users (Username)
    )`, tableName, user.Username)

	rows, err := sqlDB.Query(sqlQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func VerifyIdPass(user *pb.User) (bool, error) {
	doesUserExists, err := DoesUserExists(user.Username)

	if err != nil {
		return false, err
	}

	if !doesUserExists {
		return false, errors.New("USER DOESN'T EXISTS")
	}

	rows, err := stmtVerifyIdPass.Query(user.Username, user.Password)
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

	rows, err := stmtGetUserBySessionToken.Query(sessionToken)
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

// Return users whose name starts with searchString
func SearchUsersByPattern(searchString string) ([]string, error) {
	searchString += "%"

	rows, err := stmtSearchUsersByPattern.Query(searchString)

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

func FetchLastMessage(user string, user1 string) (*pb.Message, error) {
	tableUser := "messages_" + strings.ToLower(user)
	tableUser1 := "messages_" + strings.ToLower(user1)

	var lastMessage pb.Message
	query := `
		SELECT * FROM (
			SELECT * FROM ` + tableUser + ` WHERE Sender = ? 
			UNION ALL 
			SELECT * FROM ` + tableUser1 + ` WHERE Sender = ?
		) AS combined 
		ORDER BY DateTime DESC LIMIT 1;
	`

	rows, err := sqlDB.Query(query, user1, user)
	if err != nil {
		return &lastMessage, err
	}
	defer rows.Close()

	if rows.Next() {
		if err = rows.Scan(&lastMessage.Sender, &lastMessage.Receiver, &lastMessage.MessageContent, &lastMessage.Status, &lastMessage.DateTime); err != nil {
			return &lastMessage, err
		}
		return &lastMessage, nil
	}
	return &lastMessage, nil
}

func GetMessagesWithOffset(user string, user1 string, offset int, limit int) ([]*pb.Message, error) {
	tableUser := "messages_" + strings.ToLower(user)
	tableUser1 := "messages_" + strings.ToLower(user1)

	rows, err := sqlDB.Query("(SELECT * FROM "+tableUser+" WHERE Sender = ? UNION SELECT * FROM "+tableUser1+" WHERE Sender = ? ORDER BY DateTime DESC LIMIT ? OFFSET ?)ORDER BY DateTime;", user1, user, limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var messages []*pb.Message
	for rows.Next() {
		var message pb.Message

		if err = rows.Scan(&message.Sender, &message.Receiver, &message.MessageContent, &message.Status, &message.DateTime); err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	return messages, nil
}
