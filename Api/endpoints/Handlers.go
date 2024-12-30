package endpoints

import (
	"net/http"
	"encoding/json"
	"fmt"
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