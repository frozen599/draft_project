package helper

import (
	"strings"
)

// Replace {{PLACEHOLDER}} in string body with customer info
func substituteStr(body, title, firstName, lastName, email string) string {
	str1 := strings.ReplaceAll(body, "{{TITLE}}", title)
	str2 := strings.ReplaceAll(str1, "{{FIRST_NAME}}", firstName)
	str3 := strings.ReplaceAll(str2, "{{LAST_NAME}}", lastName)
	result := strings.ReplaceAll(str3, "{{EMAIL}}", email)

	return result
}
