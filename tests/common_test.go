package tests

import (
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
