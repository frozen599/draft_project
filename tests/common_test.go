package tests

import (
	"fmt"
	"path"
	"sync"
	"testing"

	"github.com/gopher5889/interview/helper"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	config := helper.ParseConfig("../.env")
	assert.NotEmpty(t, config)
	helper.PrintJSON(config)
}

func TestLoadCustomer(t *testing.T) {
	customers := helper.LoadCustomer("../data/customers.csv")
	assert.NotNil(t, customers)
	helper.PrintJSON(customers)
}

func TestSendMail(t *testing.T) {
	config := helper.ParseConfig("../.env")
	assert.NotNil(t, config)

	from := "example@gmail.com"
	recipients := []string{"nttuong58@gmail.com", "penguin147258369@gmail.com"}
	var wg sync.WaitGroup
	for _, r := range recipients {
		wg.Add(1)
		go helper.DoSendMail(from, r, "Test Message", config, &wg)
	}
	wg.Wait()
}

func TestProcess(t *testing.T) {
	config := helper.ParseConfig("../.env")
	assert.NotNil(t, config)

	customers := helper.LoadCustomer("../data/customers.csv")
	assert.NotNil(t, customers)
	helper.DoProcess(false, config, customers)
}

func TestFoo(t *testing.T) {
	arbitaryPath := "../output_emails//"
	result := path.Join(arbitaryPath, "emails.json")
	assert.NotEmpty(t, result)
	fmt.Println(result)
}
