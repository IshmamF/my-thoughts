package endpoints

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
	}
	fmt.Println("Connected to database")
}