package main

import (
	"log"
	"os"

	"github.com/gopher5889/interview/config"
	"github.com/gopher5889/interview/helper"
)

func main() {
	// You should run the application with syntax as follow:
	// go run main.go /path/to/email_template.json /path/to/customers.csv /path/to/output_emails /path/to/errors.csv
	if len(os.Args) != 5 {
		log.Fatal("Wrong number of arguments")
	}

	// You can also set env variables in the .env file and use helper.ParseConfig("/path/to/.env")
	config := config.NewConfig(os.Args[1], os.Args[2], os.Args[3], os.Args[4])
	helper.PrintJSON(config)
	customers := helper.LoadCustomer(config.CustomerDataFilePath)
	helper.PrintJSON(customers)
	helper.DoProcess(false, config, customers)
}
