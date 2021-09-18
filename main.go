package main

import (
	"os"

	"github.com/gopher5889/interview/config"
	"github.com/gopher5889/interview/helper"
)

func main() {
	config := config.NewConfig(os.Args[1], os.Args[2], os.Args[3], os.Args[4])
	helper.PrintJSON(config)
	customers := helper.LoadCustomer(config.CustomerDataFilePath)
	helper.PrintJSON(customers)
	helper.DoProcess(false, config, customers)
}
