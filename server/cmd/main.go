package main

import (
	"log"
	"net/http"
	"server/db"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	defer dbConn.CloseDatabase()
	http.ListenAndServe(":8080", nil)

}
