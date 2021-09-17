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

func WriteCustomer(path string, customers []*models.Customer) {

	// Filter customer that doesn't have email
	var errData []*models.Customer
	for _, cus := range customers {
		if cus.Email == "" {
			errData = append(errData, cus)
		}
	}

	// Marshal customer to write to file
	data, err := csvutil.Marshal(&errData)
	if err != nil {
		log.Fatal(err)
	}

	// if the output file does not exist, create the file
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}
