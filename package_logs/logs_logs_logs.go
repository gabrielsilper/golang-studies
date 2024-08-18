package logs

import (
	"fmt"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, v := range log {
		switch v {
		case 10071:
			return "recommendation"
		case 128269:
			return "search"
		case 9728:
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newString string
	for _, v := range log {
		if v == oldRune {
			newString = fmt.Sprintf("%s%c", newString, newRune)
			continue
		}
		newString = fmt.Sprintf("%s%c", newString, v)
	}
	return newString
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
