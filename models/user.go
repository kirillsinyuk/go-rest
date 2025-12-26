package models

import (
	"errors"
	"rest/db"
	"rest/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	insert := "INSERT INTO user (email, password) VALUES(?, ?)"
	result, err := db.DB.Exec(insert, u.Email, utils.HashPassword(u.Password))
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id

	return err
}

func (u *User) ValidateCredentials() error {
	selectUser := "SELECT id, password FROM user WHERE email = ?"
	row := db.DB.QueryRow(selectUser, u.Email)

	var hash string
	err := row.Scan(&u.ID, &hash)
	if err != nil {
		return err
	}
	pass := utils.CheckPasswordHash(u.Password, hash)

	if !pass {
		return errors.New("Invalid password")
	}

	return err
}
