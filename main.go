package main

import (
	"clean/repos"
	"clean/usecases"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	db := repos.InitTestDB()
	userUsecase := usecases.NewUserUsecase(db)

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			handler(w, r, userUsecase)
		})

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request, usecase usecases.UserUsecase) {
	if r.Method == "GET" {
		id, ok := r.URL.Query()["id"]
		if !ok || len(id) < 1 {
			fmt.Fprintf(w, "No user ID provided.")
			return
		}
		uuid, err := strconv.Atoi(id[0])
		if err != nil {
			fmt.Fprintf(w, "%v", err)
		}
		user := usecase.GetUser(uuid)
		if user.Name == "" {
			fmt.Fprintf(w, "No such user.")
			return
		}
		fmt.Fprintf(w, "%s", user.Name)
	}

	if r.Method == "POST" {
		query, ok := r.URL.Query()["name"]
		if !ok || len(query) < 1 {
			fmt.Fprintf(w, "No user name provided.")
			return
		}
		name := query[0]
		id, err := usecase.CreateUser(name)
		if err != nil {
			fmt.Fprintf(w, "Something wrong.")
			return
		}
		fmt.Fprintf(w, "User %s with ID %d has been created.", name, id)
	}
}
