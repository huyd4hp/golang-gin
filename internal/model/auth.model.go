package models

import "gorm.io/gorm"

type Auth struct {
    gorm.Model
    Email    string `gorm:"type:varchar(255);not null;unique"`
    Password string `json:"-" gorm:"type:varchar(255);not null"`
    Role     string `gorm:"type:enum('Admin','User');default:'User'"`
}
