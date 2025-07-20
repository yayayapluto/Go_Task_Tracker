package helper

import (
	"encoding/json"
	"fmt"
)

func PrintAnyAsJson(v any) {
	jsonBytes, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))
}
