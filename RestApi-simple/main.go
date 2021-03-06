package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    "json:ID"
	Name    string "json:Name"
	Content string "json:content"
}

type allTask []task

var tasks = allTask{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some Content 1",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	log.Fatal(http.ListenAndServe(":3000", router))
}
