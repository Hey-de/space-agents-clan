package main

import (
	"encoding/json"
	"os"
)

type User struct {
	Username    string
	Password    string
	Privelegies int
}
type DataBase struct {
	Users     []User
	Usercount uint
}

func ParseDB(db string) (DataBase, error) {
	var emptyDB = DataBase{}
	file, err := os.ReadFile("./db/" + db + ".json")
	if err != nil {
		return emptyDB, err
	}
	var result DataBase
	err = json.Unmarshal(file, &result)
	return result, nil
}
func UpdateCounter(db string) error {
	data, err := ParseDB(db)
	if err != nil {
		return err
	}
	data.Usercount++

	result, err := json.Marshal(data)
	os.WriteFile("./db/"+db+".json", result, 02)
	return err
}
