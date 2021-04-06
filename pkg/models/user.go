package models

import (
	"golang.org/x/crypto/bcrypt"
)

type RequestRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RequestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

//GetUser from db
func GetUser(username string) (user *User, err error) {
	var email string
	var password string
	var role  uint8

	sql := `
		SELECT email, password, role 
		FROM users 
		WHERE username = ?
	`
	req, err := DB.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer req.Close()

	err = req.QueryRow(username).Scan(&email, &password, &role)
	if err != nil {
		return nil, err
	}

	user = &User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     Role(role),
	}

	return user, nil
}

//CreateUser in db
func CreateUser(register RequestRegister) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	sql := `
		INSERT INTO users (email, username, password, role) 
		VALUES (?, ?, ?, 1)
	`

	req, err := DB.Prepare(sql)
	if err != nil {
		return err
	}
	defer req.Close()

	if _, err := req.Exec(register.Email, register.Username, hash); err != nil {
		return err
	}

	return nil
}