package main

import (
	"net/http"

	"github.com/double-coder/GoStarter.git/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
