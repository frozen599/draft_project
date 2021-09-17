package main

import (
	"github.com/gopher5889/interview/helper"
)

func main() {
	/* config := helper.ParseConfig(".env")
	fmt.Println(config.EmailTemplatePath) */

	customers := helper.LoadCustomer("./data/customers.csv")
	helper.Process("./templates/email_template.json", "./output_emails/output_emails.json", customers)
}
