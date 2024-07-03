package models
type Profile struct{
	AuthID int
	Auth Auth `gorm:"foreignKey:AuthID"`
	First_name string `gorm:"type:varchar(50)"`
	Last_name string `gorm:"type:varchar(50)"`
}