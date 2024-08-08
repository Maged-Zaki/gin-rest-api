package middlewares

import (
	"strings"

	"github.com/Maged-Zaki/gin-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func ValidateJWT(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	authFields := strings.Fields(authorization)

	if len(authFields) < 2 {
		c.AbortWithStatusJSON(401, utils.FormatResponse("Token required", nil))
		return
	}

	// Validate token
	claims, err := utils.ValidateToken(authFields[1])
	if err != nil {
		c.AbortWithStatusJSON(401, utils.FormatResponse("Invalid token", nil))
		return
	}

	userIdFloat64, ok := claims["userId"].(float64)
	if !ok {
		c.AbortWithStatusJSON(401, utils.FormatResponse("Valid token but couldn't type assert userId", nil))
	}

	userId := int64(userIdFloat64)

	email := claims["email"].(string)

	c.Set("userId", userId)
	c.Set("userEmail", email)
	// set the claims in the context
	c.Set("parsedToken", claims)
	c.Next()
}
