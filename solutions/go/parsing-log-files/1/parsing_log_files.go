package parsinglogfiles

import (
	"regexp"
)

var validLogRegex = regexp.MustCompile(`^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]`)
var splitLogLineRegex = regexp.MustCompile(`<[~*=-]*>`)
var passwordRegex = regexp.MustCompile(`".*(?i)password.*"`)
var endOfLineRegex = regexp.MustCompile(`end-of-line[0-9]*`)
var userNameRegex = regexp.MustCompile(`User\s+(\w+)`)

func IsValidLine(text string) bool {
	return validLogRegex.MatchString(text)
}

func SplitLogLine(text string) []string {
	return splitLogLineRegex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, v := range lines {
		if passwordRegex.MatchString(v) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return endOfLineRegex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	for i, line := range lines {
		splitString := userNameRegex.FindStringSubmatch(line)
		if len(splitString) < 2 {
			continue
		}
		lines[i] = "[USR] " + splitString[1] + " " + line
	}
	return lines
}
