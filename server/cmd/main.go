package main

import (
	"net/http"
	"server/router"
)

func main() {

	router.InitRouter()
	http.ListenAndServe(":8080", nil)
}
