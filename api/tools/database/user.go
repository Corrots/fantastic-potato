package database

import "log"

type User struct {
	UserId          int    `json:"user_id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (u *User) Add() error {
	stmtIn, err := db.Prepare("INSERT INTO `users`(username, password) VALUES(?, ?)")
	if err != nil {
		return err
	}
	if _, err = stmtIn.Exec(u.Username, u.Password); err != nil {
		return err
	}
	defer stmtIn.Close()
	return nil
}

func (u *User) GetByUsername() (*User, error) {
	stmtOut, err := db.Prepare("SELECT * FROM `users` WHERE username=?")
	if err != nil {
		log.Printf("db prepare err: %v\n", err)
		return nil, err
	}
	var output User
	if err := stmtOut.QueryRow(u.Username).Scan(&output.UserId, &output.Username, &output.Password); err != nil {
		log.Printf("SQL exec err: %v\n", err)
		return nil, err
	}
	defer stmtOut.Close()
	return &output, nil
}

func (u *User) Delete() error {
	stmtIn, err := db.Prepare("DELETE FROM `users` WHERE username=?")
	if err != nil {
		return err
	}
	_, err = stmtIn.Exec(u.Username)
	defer stmtIn.Close()
	return err
}
