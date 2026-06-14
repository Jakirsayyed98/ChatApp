package model

import (
	"chatapp/internal/config"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"password" db:"password"`
	Email    string    `json:"email" db:"email"`
	Status   string    `json:"status" db:"status"`
}

func InsertUser(user *User) error {
	db := config.GetDB()
	if _, err := db.Exec("Insert into users (username,password_hash,email) values ($1, $2, $3)", user.Username, user.Password, user.Email); err != nil {
		return err
	}
	return nil
}

func GetUserByMail(email string) (*User, error) {
	db := config.GetDB()
	user := &User{}

	err := db.QueryRow("SELECT id, username, password_hash, email, status FROM users WHERE email = $1", email).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Status)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByID(id string) (*User, error) {
	db := config.GetDB()
	user := &User{}
	err := db.QueryRow("SELECT id, username, email, status FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email, &user.Status)
	if err != nil {
		return nil, err
	}
	return user, nil
}
