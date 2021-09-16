package main

import (
	"fmt"

	"github.com/gopher5889/interview/helper"
)

func main() {
	config := helper.ParseConfig(".env")
	fmt.Println(config.EmailTemplatePath)

	customers := helper.LoadCustomer("./data/customers.csv")
	for _, cus := range customers {
		fmt.Println(cus.Title, cus.FirstName, cus.LastName, cus.Email)
	}
}
