package controller

import (
	"basic/common"
	"basic/global"
	"basic/internal/schema"
	"basic/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth interface {
	SignUp(c *gin.Context)
	LogIn(c *gin.Context)
}

type auth struct {
}

func AuthController() Auth {
	return &auth{}
}

func (ac *auth) SignUp(c *gin.Context) {
	authService := service.AuthService(global.GetDB())
	var data schema.SignUp
	if err := c.ShouldBindJSON(&data); err != nil {
		error := common.BadRequestError(err, "Invalid Input Error", err.Error(), "BadRequest")
		c.JSON(error.StatusCode, error)
		return
	}
	// Kiểm tra tính hợp lệ của dữ liệu đăng ký
	if errInput := data.Validation(); errInput != nil {
		error := common.BadRequestError(errInput, "Invalid Input Error", errInput.Error(), "BadRequest")
		c.JSON(error.StatusCode, error)
		return
	}
	// Check Data
	_, errFind := authService.FindOne(data.Email)
	if errFind == nil{
		error := common.ConflictError(errFind,"Email has been used","","ConflictRequest")
		c.JSON(error.StatusCode,error)
		return
	}
	// Not Found - Insert Data
	_, errCreate := authService.Create(data)
	if errCreate != nil{
		error := common.DB_ERROR(errCreate)
		c.JSON(error.StatusCode,error)
		return
	}
	c.JSON(
		http.StatusCreated,
		gin.H{
			"status_code":201,
			"message":"Sign Up Successully",
		},
	);
}
func (ac *auth) LogIn(c *gin.Context){
	authService := service.AuthService(global.GetDB())
	var data schema.LogIn
	if err := c.ShouldBindJSON(&data); err != nil {
		error := common.BadRequestError(err, "Invalid Input Error", err.Error(), "BadRequest")
		c.JSON(error.StatusCode, error)
		return
	}
	// Kiểm tra tính hợp lệ của dữ liệu đăng nhập
	if errInput := data.Validation(); errInput != nil {
		error := common.BadRequestError(errInput, "Invalid Input Error", errInput.Error(), "BadRequest")
		c.JSON(error.StatusCode, error)
		return
	}
	// Check Data
	account, errFind := authService.FindOne(data.Email)
	if errFind != nil{
		error := common.NotFoundError(errFind,"Email Address is not Registered",errFind.Error(),"NotFoundRequest")
		c.JSON(error.StatusCode,error)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"data":account,
		},
	);

}
