package helper

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gopher5889/interview/config"
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

func ProcessEmailPerCustomer(templatePath string, customer *models.Customer, ch chan models.Email) {
	// Check if path to email_template.json exists
	if _, err := os.Stat(templatePath); err != nil {
		log.Fatal(err)
	}

	// Read template file then unmarshal it to EmailTemplate
	bytes, err := os.ReadFile(templatePath)
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
		email := models.Email{Template: &template, To: customer.Email}
		ch <- email
	}
}

func DoWriteOutput(outputPath string, result []models.Email) {
	// Marshal json and then write to output file
	var temp []models.EmailTemplate
	for _, email := range result {
		temp = append(temp, *email.Template)
	}

	data, err := json.MarshalIndent(&temp, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	// if the output file does not exist, create the file
	file, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Because the "<" is chaned to \u003c during serialization
	written := strings.ReplaceAll(string(data), "\\u003c", "<")
	_, err = file.WriteString(written)
	if err != nil {
		log.Fatal(err)
	}

}

func DoProcess(sendMail bool, config *config.Config, customers []*models.Customer) {
	ch := make(chan models.Email)

	go func() {
		for _, cus := range customers {
			ProcessEmailPerCustomer(config.EmailTemplatePath, cus, ch)
		}
		// signal that no more data will be sent through the channel
		close(ch)
	}()

	var result []models.Email
	// Loop until the channel is closed, this is a pipeline model
	for tmpl := range ch {
		result = append(result, tmpl)
	}

	if !sendMail {
		DoWriteOutput(config.OutputEmailPath, result)
	} else {
		if config.SmtpHost == "" || config.SmtpPort == "" || config.Password == "" {
			log.Fatal("SendMail config vars must be set")
		}

		var wg sync.WaitGroup
		for _, email := range result {
			wg.Add(1)
			go DoSendMail(email.Template.From, email.To, email.Template.Body, config, &wg)
		}
		wg.Wait()
	}

	WriteErrCustomer(config.ErrorPath, customers)
}
