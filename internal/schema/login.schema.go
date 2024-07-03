package schema

import (
	"errors"
	"net/mail"
)

type LogIn struct {
    Email    string 
    Password string 
}
func (LogIn) TableName() string {
    return "auths"
}

func (data *LogIn) Validation() error {
    // Check Email Input
    if data.Email == ""{
        return errors.New("email required")
    }
    // Check Email Format
    _, err := mail.ParseAddress(data.Email)
    if err != nil {
        return errors.New("invalid email format")
    }
    // Check Password Input
    if data.Password == ""{
        return errors.New("password required")
    }
    return nil
}

