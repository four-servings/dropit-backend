package main

import (
	_ "github/four-servings/dropit-backend/stuff"
	_ "github/four-servings/dropit-backend/user"
	"net/http"
)

func main() {
	http.ListenAndServe(":5000", nil)
}
