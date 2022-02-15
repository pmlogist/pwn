package main

import (
	"fmt"
	"log"
)

type User struct {
	Id       string
	Name     string
	Username string
	Password string
	IsAdmin  bool
	Sex      string
	Likes    int
	Token    string
}

func (r *Router) GetAll() ([]User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return []User{}, err
	}

	defer rows.Close()

	var users []User
	var user User
	for rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Username,
			&user.Password,
			&user.Sex,
			&user.Likes,
			&user.IsAdmin,
		)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, User{
			Id:       user.Id,
			Name:     user.Name,
			Username: user.Username,
			Password: user.Password,
			Sex:      user.Sex,
			Likes:    user.Likes,
			IsAdmin:  user.IsAdmin,
		})

	}

	return users, nil
}

func (r *Router) GetOneByUsername(username string) (User, error) {
	rows, err := r.db.Query("SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		return User{}, err
	}

	defer rows.Close()

	var user User
	for rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Username,
			&user.Password,
			&user.Sex,
			&user.Likes,
			&user.IsAdmin,
		)
		if err != nil {
			log.Fatal(err)
		}

	}

	return user, nil
}

func (r *Router) GetOneById(id string) (User, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE id= %s;", id)
	rows, err := r.db.Query(query)
	if err != nil {
		log.Print(err)
		return User{}, err
	}

	defer rows.Close()

	var user User
	for rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Username,
			&user.Password,
			&user.Sex,
			&user.Likes,
			&user.IsAdmin,
		)
		if err != nil {
			log.Fatal(err)
		}

	}

	return user, nil
}

func (r *Router) DeleteOneById(id string) {
	stmt, _ := r.db.Prepare("UPDATE users SET is_admin = true WHERE id = 2;")
	stmt.Exec()
}

func (r *Router) LikeOneById(id string) {
	query := "UPDATE users SET likes = likes + 1 WHERE id=" + id
	//query := "UPDATE users SET is_admin=true WHERE id=2; UDATE users SET likes=likes+1 WHERE id=2;"
	fmt.Println(query)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Print(err)
	}
	stmt.Exec()
}
