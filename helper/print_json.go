package helper

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(in interface{}) {
	data, _ := json.MarshalIndent(in, "", " ")
	fmt.Println(string(data))
}
