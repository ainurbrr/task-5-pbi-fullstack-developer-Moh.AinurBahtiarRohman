package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"task-5-vix-btpns-Moh.AinurBahtiarRohman/helpers"
	"task-5-vix-btpns-Moh.AinurBahtiarRohman/models"
)

func AuthMiddleware(db *gorm.DB) echo.HandlerFunc {
	return func(c *echo.Context) {
		var user models.User
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helpers.ApiResponse(http.StatusUnauthorized, "error", nil, "Unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//split the token and take the value
		tokenString := ""
		dataToken := strings.Split(authHeader, " ")
		if len(dataToken) == 2 {
			tokenString = dataToken[1]
		}

		token, err := helpers.ValidateToken(tokenString)
		if err != nil {
			response := helpers.ApiResponse(http.StatusUnauthorized, "error", nil, "Unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helpers.ApiResponse(http.StatusUnauthorized, "error", nil, "Unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		err = db.First(&user, userID).Error
		if err != nil {
			response := helpers.ApiResponse(http.StatusUnauthorized, "error", nil, "Unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
