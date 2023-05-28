package stringUtil

import "fmt"

func CreateFormattedString(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}