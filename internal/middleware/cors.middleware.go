package middleware

import (
	"basic/common"
	"slices"

	"github.com/gin-gonic/gin"
)

func CORS(allowedOrigin[]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if !slices.Contains(allowedOrigin,origin) && !slices.Contains(allowedOrigin,"*"){
			err := common.ForbiddenError(nil,"Origin not allowed","Origin not allowed","BadRequest");
			c.JSON(err.StatusCode,err);
			c.Abort()
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Next()
	}
}
