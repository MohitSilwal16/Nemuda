package utils

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"unicode"
)

func ClearScreen() {
	var cmd *exec.Cmd

	// Check the operating system to determine the appropriate clear command
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear") // for Unix-like systems
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls") // for Windows
	default:
		log.Println("Unsupported platform.")
		return
	}

	// Execute the clear command
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func IsPasswordInFormat(s string) bool {
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) < 8 || len(s) > 20 {
		return false
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSpecial
}

func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func TrimGrpcErrorMessage(errMsg string) string {
	// Split the error message
	parts := strings.Split(errMsg, "desc = ")
	if len(parts) > 1 {
		// Return the part after "desc = "
		return parts[1]
	}
	// Return the original error message if "desc = " is not found
	return errMsg
}

func ReturnAlertMessage(msg string) string {
	return "<script>alert('" + TrimGrpcErrorMessage(msg) + "')</script>"
}

// FileHeaderToBytes converts a *multipart.FileHeader to a byte slice.
func FileHeaderToBytes(fileHeader *multipart.FileHeader) ([]byte, error) {
	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content into a buffer
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
