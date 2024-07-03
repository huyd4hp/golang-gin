package service

import (
	models "basic/internal/model"
	"basic/internal/schema"

	"gorm.io/gorm"
)

type Auth interface {
    FindOne(email string) (models.Auth, error)
    Create(data schema.SignUp) (int,error)
}

type auth struct {
    db *gorm.DB
}

func AuthService(db *gorm.DB) Auth {
    return &auth{db: db}
}

func (as *auth) FindOne(email string) (models.Auth, error) {
    var auth models.Auth
    result := as.db.Where("email = ?", email).First(&auth)
    if result.Error != nil {
        return auth, result.Error
    }
    return auth, nil
}

func (as *auth) Create(data schema.SignUp) (int,error) {
    result := as.db.Create(&data);
    if result.Error != nil{
        return 0,result.Error;
    }
    return data.ID,nil
}
