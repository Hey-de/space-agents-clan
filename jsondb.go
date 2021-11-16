package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Username    string
	Password    string
	Privelegies int
}
type DataBase struct {
	Users     map[string]User
	Usercount uint
}

func (DataBase) checkPassword(database string, login string, password string) (result bool, er error) {
	var encryptedPasswd = sha1.Sum([]byte(password))
	var encryptedPasswdDB = fmt.Sprintf("% x", encryptedPasswd)
	var db, err = ParseDB(database)
	if err != nil {
		return false, err
	}
	if db.Users[login].Password == encryptedPasswdDB {

	}
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
