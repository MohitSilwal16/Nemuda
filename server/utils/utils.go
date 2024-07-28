package utils

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"time"
	"unicode"

	"google.golang.org/grpc"
)

// ANSI escape codes for background colors
const (
	ColorReset      = "\033[0m"
	BackgroundGreen = "\033[42m" // Success
	BackgroundRed   = "\033[41m" // Error
	BackgroundBlue  = "\033[44m" // Method
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

func TokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
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

func IsMessageMalicious(message string) bool {
	// List of potentially dangerous patterns
	dangerousPatterns := []string{
		`(?i)<script.*?>`,
		`(?i)javascript:`,
		`(?i)on\w+\s*=`,
		`(?i)data:`,
		`(?i)vbscript:`,
		`(?i)<iframe`,
		`(?i)<embed`,
		`(?i)<object`,
		`(?i)<img.*?onerror`,
		`(?i)<[^>]*on\w+\s*=`,
		`(?i)data:[^,]*;base64,`,
		`(?i)<marquee`,
		`(?i)<blink`,
		`(?i)<svg`,
		`(?i)<xml`,
		`<[^>]*>`,
	}

	// Check for dangerous patterns
	for _, pattern := range dangerousPatterns {
		if matched, _ := regexp.MatchString(pattern, message); matched {
			return true
		}
	}
	return false
}

func BytesToMultipartFile(fileBytes []byte, fileName string) multipart.File {
	return &memoryFile{
		name:   fileName,
		reader: bytes.NewReader(fileBytes),
	}
}

type memoryFile struct {
	name   string
	reader *bytes.Reader
}

func (f *memoryFile) Read(p []byte) (n int, err error) {
	return f.reader.Read(p)
}

func (f *memoryFile) Close() error {
	return nil
}

func (f *memoryFile) Seek(offset int64, whence int) (int64, error) {
	return f.reader.Seek(offset, whence)
}

// Add the ReadAt method to satisfy the multipart.File interface
func (f *memoryFile) ReadAt(p []byte, off int64) (n int, err error) {
	return f.reader.ReadAt(p, off)
}

// StructuredLoggerInterceptor logs requests in a format similar to Gin Gonic’s logger
func StructuredLoggerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Log start time and method
		start := time.Now()
		method := info.FullMethod

		// Handle the request
		resp, err := handler(ctx, req)

		// Log response status and duration with background color coding
		duration := time.Since(start).Milliseconds()
		status := "Success"            // Default status
		statusColor := BackgroundGreen // Default color
		if err != nil {
			status = "Failed "
			statusColor = BackgroundRed
		}

		// Format log output
		methodField := fmt.Sprintf("%-40s", method)      // Align method column
		statusField := fmt.Sprintf("%-7s", status)       // Align status column
		durationField := fmt.Sprintf("%10dms", duration) // Align duration column

		log.Printf("%s %s %s | %s |%s %s %s",
			statusColor,
			statusField,
			ColorReset,
			durationField,
			BackgroundBlue,
			methodField,
			ColorReset,
		)

		return resp, err
	}
}
