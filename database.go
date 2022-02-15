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
		"is_admin" BOOLEAN
	  );`

	log.Println("Creating users table...")
	stmt, err := db.Prepare(createUsersTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
	log.Println("users table created")
}

func InsertUser(db *sql.DB, name string, username string, password string, sex string, likes int, isAdmin bool) {
	log.Println("Inserting users record ...")
	insertUserSQL := `INSERT INTO users(name, username, password, sex, likes, is_admin) VALUES (?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertUserSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(name, username, password, sex, likes, isAdmin)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func InitDatabase(db *sql.DB) {
	CreateDatabase(db)

	InsertUser(db, "Jane", "jane", "192837465", "F", 10, true)
	InsertUser(db, "Petra", "petra", "password", "F", 12, false)
	InsertUser(db, "Bob", "bob", "password", "M", 11, false)
}
