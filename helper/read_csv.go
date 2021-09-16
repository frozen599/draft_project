package helper

import (
	"log"
	"os"

	"github.com/gopher5889/interview/models"
	"github.com/jszwec/csvutil"
)

func LoadCustomer(path string) []*models.Customer {
	_, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var result []*models.Customer
	err = csvutil.Unmarshal(bytes, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
