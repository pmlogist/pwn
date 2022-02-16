package main

import (
	"fmt"
	"log"
)

func (r *Router) GetAll() ([]User, error) {
	rows, err := r.db.Query("SELECT * FROM users WHERE is_admin = false;")
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
			&user.Avatar,
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
			Avatar:   user.Avatar,
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
			&user.Avatar,
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
			&user.Avatar,
		)
		if err != nil {
			log.Fatal(err)
		}

	}

	return user, nil
}

func (r *Router) DeleteOneById(id string) {
	// Unsafe
	query := fmt.Sprintf("DELETE FROM users WHERE id = '%s'", id)
	stmt, _ := r.db.Prepare(query)
	stmt.Exec()
}

func (r *Router) LikeOneById(id string) {
	query := fmt.Sprintf("UPDATE users SET likes = likes + 1 WHERE id = %s", id)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Print(err)
	}
	stmt.Exec()
}
