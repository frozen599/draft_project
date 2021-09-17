package main

import (
	"github.com/gopher5889/interview/helper"
)

func main() {
	config := helper.ParseConfig(".env")
	helper.PrintJSON(config)
	customers := helper.LoadCustomer("./data/customers.csv")
	helper.PrintJSON(customers)
	helper.DoProcess(false, config, customers)
}
