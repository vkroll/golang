package userdb

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password []byte
}

var users = map[string]User{}

func CreateUser(u string, p []byte) (bool, error) {
	if UserExists(u) {
		return false, fmt.Errorf("User %s already exists", u)
	}
	password, err := bcrypt.GenerateFromPassword(p, 8)
	if err != nil {
		panic(err)
	}
	user := User{Username: u, Password: password}
	users[u] = user
	//log.Print(users)
	return true, nil
}

func UserExists(u string) bool {
	_, ok := users[u]
	return ok
}

func ValidateUser(u string, p []byte) bool {
	U, ok := users[u]
	if !ok {
		return false
	}

	err := bcrypt.CompareHashAndPassword(U.Password, p)
	if err != nil {
		log.Printf("Login failed for %s", u)
		return false
	}
	return true

}

func LoadUsers() {
	var rdr io.Reader
	f, err := os.Open("users.json")
	if err != nil {
		rdr = strings.NewReader("{}")
	} else {
		defer f.Close()

		rdr = f
	}
	err = json.NewDecoder(rdr).Decode(&users)
	if err != nil {
		panic(err)
	}
	for _, u := range users {
		users[u.Username] = u
	}
}

func SaveUsers() {
	f, err := os.Create("users.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(users)
	if err != nil {
		panic(err)
	}
}
