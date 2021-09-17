package tests

import (
	"sync"
	"testing"

	"github.com/gopher5889/interview/helper"
	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
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
