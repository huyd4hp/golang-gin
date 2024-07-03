package initialize

import (
	"basic/internal/controller"
	"basic/internal/middleware"

	"github.com/gin-gonic/gin"
)
func InitRouter()*gin.Engine {
	// AlowedOrigin
	allowedOrigins := []string{"*"}
	r:= gin.Default()
	r.Use(middleware.CORS(allowedOrigins))
	r.NoRoute(middleware.ApiUndefined())
	r.NoMethod(middleware.MethodNotAllowed())

	v1 := r.Group("/api/v1")
	{
		authController := controller.AuthController()
		auth := v1.Group("/auth")
		{
			auth.POST("/signup",authController.SignUp)
			auth.POST("/login",authController.LogIn)
		}
	}
	return r;
}