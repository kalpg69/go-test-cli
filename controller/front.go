package controller

import "net/http"

func RegisterUserController() {
	uc := newuserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)

}
