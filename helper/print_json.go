package helper

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

func PrintJSON(in interface{}) {
	data, _ := jsoniter.MarshalIndent(in, " ", "    ")
	fmt.Println(data)
}
