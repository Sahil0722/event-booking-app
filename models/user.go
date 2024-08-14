package models

import (
	"errors"

	"github.com/Sahil0722/event-booking-app/db"
	"github.com/Sahil0722/event-booking-app/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT into users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, hashPassword)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	u.ID = id

	return err
}

func (u *User) AuthenticateUser() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retreivedPassword string
	err := row.Scan(&u.ID, &retreivedPassword)

	if err != nil {
		return errors.New("Invalid Credentials")
	}

	isPasswordValid := utils.ValidatePassword(u.Password, retreivedPassword)

	if !isPasswordValid {
		return errors.New("Invalid Credentials")
	}

	return nil
}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
