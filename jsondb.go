package main

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"io/ioutil"
)

func findUser(login string) (bool, error) {
	var db, err = ParseDB("users")
	if err != nil {
		return false, err
	}
	_, ok := db[login]
	fmt.Println(db)
	return ok, nil
}

func createPost(login string, password interface{}, title string, content string, picture []string) error {
	fmt.Println(login)
	fmt.Printf("%x", password.([20]uint8))
	ok, err := comparePasswords(login, fmt.Sprintf("%x", password.([20]uint8)))
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("authefication failed")
	}
	db, err := ParseDBSlice("data")
	if err != nil {
		return err
	}
	db["Notes"] = append(db["Notes"], map[string]interface{}{
		"Owner":   login,
		"Title":   title,
		"Content": content,
		"Image":   picture,
	})
	err = UpdateDB(db, "data")
	return err
}
func register(login string, password string) error {
	oldUser, err := findUser(login)
	if err != nil {
		return err
	}
	if oldUser {
		return errors.New("user already exists")
	}
	db, err := ParseDB("users")
	if err != nil {
		return err
	}
	sum := sha1.Sum([]byte(password))
	db[login] = fmt.Sprintf("%x", sum)
	err = UpdateDB(db, "users")
	return err
}
func comparePasswords(login string, password string) (bool, error) {
	var db, err = ParseDB("users")
	if err != nil {
		return false, err
	}
	if db[login] == password {
		return true, nil
	} else {
		return false, nil
	}
}
func checkPassword(login string, password string) (bool, error) {
	var encryptedPasswd = sha1.Sum([]byte(password))
	var encryptedPasswdDB = fmt.Sprintf("%x", encryptedPasswd)
	result, err := comparePasswords(login, encryptedPasswdDB)
	return result, err
}
func ParseDBInt(db string) (map[string]int, error) {
	file, err := os.ReadFile("./db/" + db + ".json")
	if err != nil {
		return nil, err
	}
	var result map[string]int
	err = json.Unmarshal(file, &result)
	return result, err
}
func ParseDBSlice(db string) (map[string][]interface{}, error) {
	file, err := os.ReadFile("./db/" + db + ".json")
	if err != nil {
		return nil, err
	}
	var result map[string][]interface{}
	err = json.Unmarshal(file, &result)
	return result, err
}

func ParseDB(db string) (map[string]interface{}, error) {
	file, err := os.Open("./db/" + db + ".json")
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	return result, err
}
func UpdateDB(data interface{}, db string) error {
	mu.Lock()
	defer mu.Unlock()
	result, err := json.Marshal(data)
	if err != nil {
		return err
	}
	file, err := os.Open("./db/"+db+".json")
	if err != nil {
		return err
	}
	_, err = os.Write(result)
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
