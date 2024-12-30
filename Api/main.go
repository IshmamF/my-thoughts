package main

import (
	"fmt"
	"github.com/IshmamF/my-thoughts/Api/database"
	"net/http"
	//"github.com/IshmamF/my-thoughts/Api/endpoints"
	"encoding/json"
)

var (
	db = &database.DB{}
)

func ReadThoughts (w http.ResponseWriter, r *http.Request) {
	thoughts, err := db.GetRows()
	if err != nil {
		fmt.Println("Error with reading thoughts from db")
	}
	jsonData, err := json.Marshal(thoughts)
	if err != nil {
		fmt.Println("Error with converting thoughts to json")
	}
	w.Write(jsonData)
}


func main() {
	db.Connect()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /read-thoughts", ReadThoughts)
	http.ListenAndServe(":5000", mux)
	

	/*
	err = db.InsertRow("Hello, World!", "Thought")
	if err != nil {
		fmt.Println("Error inserting row", err)
	} else {
		fmt.Println("Row inserted")
	} */
}