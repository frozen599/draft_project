package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gopher5889/interview/models"
)

// Replace {{PLACEHOLDER}} in string body with customer info
func substituteStr(body, title, firstName, lastName string) string {
	str1 := strings.ReplaceAll(body, "{{TITLE}}", title)
	str2 := strings.ReplaceAll(str1, "{{FIRST_NAME}}", firstName)
	str3 := strings.ReplaceAll(str2, "{{LAST_NAME}}", lastName)
	result := strings.ReplaceAll(str3, "{{TODAY}}", time.Now().Format("January 2, 2006"))

	return result
}

func ProcessEmailPerCustomer(path string, customer *models.Customer, ch chan models.EmailTemplate) {
	// Check if path to email_template.json exists
	if _, err := os.Stat(path); err != nil {
		log.Fatal(err)
	}

	// Read template file then unmarshal it to EmailTemplate
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var template models.EmailTemplate
	err = json.Unmarshal(bytes, &template)
	if err != nil {
		log.Fatal(err)
	}

	if customer.Email != "" {
		newBody := substituteStr(template.Body, customer.Title, customer.FirstName, customer.LastName)
		template.Body = newBody
		ch <- template
	}
}

func Process(templatePath, outPath string, customers []*models.Customer) {
	ch := make(chan models.EmailTemplate)

	go func() {
		for _, cus := range customers {
			ProcessEmailPerCustomer(templatePath, cus, ch)
		}
		close(ch)
	}()

	var result []models.EmailTemplate
	for tmpl := range ch {
		fmt.Println(tmpl.Subject, tmpl.Body, tmpl.From)
		result = append(result, tmpl)
	}

	// Marshal json and then write to output file
	data, err := json.MarshalIndent(&result, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	// if the output file does not exist, create the file
	file, err := os.OpenFile(outPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//
	written := strings.ReplaceAll(string(data), "\\u003c", "<")
	fmt.Println(written)
	_, err = file.WriteString(written)
	if err != nil {
		log.Fatal(err)
	}
}
