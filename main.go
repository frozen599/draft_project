package main

import (
	"log"
	"os"

	"github.com/gopher5889/interview/config"
	"github.com/gopher5889/interview/helper"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatal("Wrong number of arguments")
	}

	config := config.NewConfig(os.Args[1], os.Args[2], os.Args[3], os.Args[4])
	helper.PrintJSON(config)
	customers := helper.LoadCustomer(config.CustomerDataFilePath)
	helper.PrintJSON(customers)
	helper.DoProcess(false, config, customers)
}
