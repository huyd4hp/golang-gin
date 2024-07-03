package middleware

import (
	"basic/common"

	"github.com/gin-gonic/gin"
)

func MethodNotAllowed() gin.HandlerFunc{
	return func(c *gin.Context) {
		err := common.MethodNotAllowedError(nil,"Method Not Allowed","Method Not Allowed","BadRequest")
		c.JSON(
			err.StatusCode,
			err,
		)    
		c.Abort()   
    }
}