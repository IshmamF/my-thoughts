package main

import (
	"fmt"
	"github.com/IshmamF/my-thoughts/Api/database"
)


var (
	db = &database.DB{}
)

func main() {
	db.Connect()
	err := db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database")
	} else {
		fmt.Println("Connected to database")
	}

	// Insert a row
	err = db.InsertRow("Hello, World!", "Thought")
	if err != nil {
		fmt.Println("Error inserting row", err)
	} else {
		fmt.Println("Row inserted")
	}

	// Get all rows
	thoughts, err := db.GetRows()
	if err != nil {
		fmt.Println("Error getting rows", err)
	} else {
		for _, thought := range thoughts {
			fmt.Println(thought)
		}
	}
}