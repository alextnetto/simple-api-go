package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/alextnetto/pkg/entity"
	"github.com/alextnetto/pkg/repository"
)

func HandleNewUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, err.Error())
	}

	newUser := entity.User{}

	err = json.Unmarshal(body, &newUser)
	if err != nil {
		io.WriteString(w, err.Error())
	}

	// TODO: connect with database and store

	result, err := repository.NewUserRepository()

	res, err := json.MarshalIndent(newUser, "", " ")
	if err != nil {
		io.WriteString(w, err.Error())
	}

	io.WriteString(w, string(res))
}

func main() {
	http.HandleFunc("/", HandleNewUser)
	http.ListenAndServe(":80", nil)
}
