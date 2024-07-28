package utils

import (
	"regexp"
)

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
