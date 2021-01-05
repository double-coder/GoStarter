package main

import (
	"fmt"

	"github.com/double-coder/GoStarter.git/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Aditya",
		LastName:  "Pathania",
	}
	fmt.Println(u)
}
