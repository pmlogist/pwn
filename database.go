package main

import (
	"database/sql"
	"log"
)

func CreateDatabase(db *sql.DB) {
	createUsersTableSQL := `CREATE TABLE users (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" VARCHAR(60),
		"username" VARCHAR(60),
		"password" VARCHAR(225),
		"sex" VARCHAR(1),
    "likes" INTEGER,
		"is_admin" BOOLEAN,
		"avatar" VARCHAR(20)
	  );`

	log.Println("Creating users table...")
	stmt, err := db.Prepare(createUsersTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
	log.Println("users table created")
}

func InsertUser(db *sql.DB, name string, username string, password string, sex string, likes int, isAdmin bool, avatar string) {
	log.Println("Inserting users record ...")
	insertUserSQL := `INSERT INTO users(name, username, password, sex, likes, is_admin, avatar) VALUES (?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertUserSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(name, username, password, sex, likes, isAdmin, avatar)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func InitDatabase(db *sql.DB) {
	CreateDatabase(db)

	InsertUser(db, "Jane", "jane", "password", "F", 10, true, "&#128103;&#127998;")
	InsertUser(db, "Petra", "petra", "password", "F", 12, false, "&#128105;&#127997;")
	InsertUser(db, "Bob", "bob", "password", "M", 11, false, "&#128102;&#127997;")
	InsertUser(db, "Marta", "marta", "password", "M", 11, false, "&#128120;&#127996")
	InsertUser(db, "John", "john", "password", "M", 11, false, "&#128102;&#127999;")
	InsertUser(db, "Evgenia", "evgenia", "password", "M", 11, false, "&#128103;&#127995;")
	InsertUser(db, "Mark", "mark", "password", "M", 11, false, "&#128104;&#127995;")
}
