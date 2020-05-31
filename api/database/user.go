package database

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

func (u *User) GetByUsername() error {
	stmtOut, err := db.Prepare("SELECT password FROM `users` WHERE username=?")
	if err != nil {
		return err
	}
	if err := stmtOut.QueryRow(u.Username).Scan(&u.Password); err != nil {
		// If the query selects no rows, the *Row's Scan will return ErrNoRows.
		return err
	}
	defer stmtOut.Close()
	return nil
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
