package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

const port = "8080"

func main() {
	os.Remove("./database.db")
	router := Router{}

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	router.db = db

	InitDatabase(db)

	r := httprouter.New()

	r.GET("/", router.Index)
	r.POST("/", router.Login)
	r.GET("/profile", router.Profile)
	r.POST("/profile", router.ProfileLikesAndDelete)
	r.GET("/nice-try", router.NiceTry)

	log.Printf("Server started on: http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
