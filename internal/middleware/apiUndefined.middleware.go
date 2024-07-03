package middleware

import (
	"basic/common"

	"github.com/gin-gonic/gin"
)

func ApiUndefined() gin.HandlerFunc {
    return func(c *gin.Context) {
		err := common.NotFoundError(nil,"Not Found","ApiUndefined","BadRequest")
		c.JSON(
			err.StatusCode,
			err,
		)
		c.Abort()
    }
}