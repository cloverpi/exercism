package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	var logTypes = map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}
	for _, v := range log {
		t, ok := logTypes[v]
		if ok {
			return t
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newLog := make([]rune, 0, len(log))
	for _, v := range log {
		thisRune := v
		if thisRune == oldRune {
			thisRune = newRune
		}
		newLog = append(newLog, thisRune)
	}
	return string(newLog)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
