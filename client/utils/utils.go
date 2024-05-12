package utils

import (
	"log"
	"os"
	"os/exec"
	"runtime"
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
