package main

import (
	"encoding/json"
	"os"
)

func ParseDB(db string) (interface{}, error) {
	file, err := os.ReadFile("./db/" + db + ".json")
	if err != nil {
		return nil, err
	}
	var result interface{}
	err = json.Unmarshal(file, &result)
	return result, nil
}
