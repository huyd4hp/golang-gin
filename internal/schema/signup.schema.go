package schema

import (
	"errors"
	"net/mail"
)

type SignUp struct {
    ID       int `json:"-" gorm:"column:id;"`
    Email    string `json:"email" gorm:"column:email;"`
    Password string `json:"password" gorm:"column:password;"`
    Role     string `json:"role" gorm:"column:role;"`
}
func (SignUp) TableName() string {
    return "auths"
}
func (data *SignUp) Validation() error {
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
    // Check Register Role
    if data.Role == ""{
        data.Role = "User"
    }
    return nil
}