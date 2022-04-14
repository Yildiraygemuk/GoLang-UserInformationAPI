package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	model "GoWork/models"
)

func main() {
	apiRoot := "/api"

	http.HandleFunc(apiRoot+"/Users", func(w http.ResponseWriter, r *http.Request) {
		message := loadUsers()
		output, err := json.Marshal(message)
		checkError(err)

		fmt.Fprintf(w, string(output))
	})
	http.ListenAndServe(":1905", nil)
}
func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error: ", err.Error())
		os.Exit(1)
	}
}

func loadUsers() []model.User {
	bytes, err := ioutil.ReadFile("json/users.json")
	checkError(err)
	var users []model.User
	json.Unmarshal(bytes, &users)
	return users
}

func loadInterests() []model.Interest {
	bytes, err := ioutil.ReadFile("json/interest.json")
	checkError(err)
	var interests []model.Interest
	json.Unmarshal(bytes, &interests)
	return interests
}
