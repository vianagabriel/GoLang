package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type User struct {
	ID   int
	Name string
	Age  int
}

var Users = []User{
	{ID: 1, Name: "Gabriel", Age: 25},
	{ID: 2, Name: "Rebeca", Age: 23},
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", addUser)
	mux.HandleFunc("/users", listUsers)
	mux.HandleFunc("/user", listUser)
	mux.HandleFunc("/delete", deleteUser)

	fmt.Println("Server is running is port 8080")
	http.ListenAndServe(":8080", mux)
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json") // setando o tipo de resposta como json

	body, err := io.ReadAll(r.Body) // lendo o que está sendo passado no corpo da requisição
	if err != nil {
		panic(err)
	}

	var newUser User // novo usuário

	json.Unmarshal(body, &newUser)     // fazendo um unmarshal do que foi passado no corpo da requisição e salvando em newUser
	Users = append(Users, newUser)     // Adicionando o novo usuário ao slice Users
	json.NewEncoder(w).Encode(newUser) // transformando a resposta que será mostrada ao usuário para json
}

func listUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json") // setando o tipo de resposta como json

	IdParam := r.URL.Query().Get("id")
	parseID, _ := strconv.Atoi(IdParam)

	for _, User := range Users {
		if User.ID == parseID {
			json.NewEncoder(w).Encode(User)
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json") // setando o tipo de resposta como json

	IdParam := r.URL.Query().Get("id")
	parseID, _ := strconv.Atoi(IdParam)
	newSlice := []User{}

	for _, User := range Users {
		if User.ID != parseID {
			newSlice = append(newSlice, User)
		}
	}

	Users = newSlice

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Users)
}
