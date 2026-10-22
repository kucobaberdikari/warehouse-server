package middleware

import (
	"management-warehouse/helper"
	"management-warehouse/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("currentUser")
		currentUser, ok := value.(user.User)
		if !exists || !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil))
			return
		}

		for _, role := range allowedRoles {
			if currentUser.Role == role {
				c.Next()
				return
			}
		}
		Meta := gin.H{
			"message": "You not authorize to access this API",
		}

		c.AbortWithStatusJSON(http.StatusForbidden,
			helper.APIResponse("Forbidden", http.StatusForbidden, "error", Meta))
	}
}
