package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"os"
)

type Users struct {
	Users map[string]string
}

type Counter struct {
	Usercount int
}

func checkPassword(login string, password string) (bool, error) {
	var encryptedPasswd = sha1.Sum([]byte(password))
	var encryptedPasswdDB = fmt.Sprintf("% x", encryptedPasswd)
	var db, err = ParseDB("users")
	if err != nil {
		return false, err
	}
	if db[login] == encryptedPasswdDB {
		return true, nil
	} else {
		return false, nil
	}
}
func ParseDBInt(db string) (map[string]int, error) {
	file, err := os.ReadFile("./db/" + db + ".json")
	if err != nil {
		return nil, err
	}
	var result map[string]int
	err = json.Unmarshal(file, &result)
	return result, nil
}
func ParseDB(db string) (map[string]interface{}, error) {
	file, err := os.ReadFile("./db/" + db + ".json")
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(file, &result)
	return result, nil
}
func UpdateDB(data interface{}, db string) error {
	result, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile("./db/"+db+".json", result, 02)
	return err
}
func UpdateCounter() error {
	data, err := ParseDBInt("index")
	if err != nil {
		return err
	}
	data["Usercount"]++
	err = UpdateDB(data, "index")
	return err
}
