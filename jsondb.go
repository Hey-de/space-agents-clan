package main

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"os"
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