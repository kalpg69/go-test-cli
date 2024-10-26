package main

import (
	"net/http"

	"github.com/kalpg69/go-test-cli/controller"
)

func main() {

	controller.RegisterUserController()
	http.ListenAndServe(":3000", nil)
}
