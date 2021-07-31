package middlewares

import (
	"net/http"
	"parky/client/token"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "You need to be authorized to access this page")
			c.Abort()
			return
		}
		data, err1 := token.ExtractTokenAuth(c.Request)
		if err1 != nil {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "not authorized",
				})
			return
		}
		c.Set("UserID", data.UserID)
		c.Next()
	}
}
