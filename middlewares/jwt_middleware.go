package middlewares

import (
	"net/http"
	"strings"

	"github.com/ainurbrr/task-5-vix-btpns-Moh.AinurBahtiarRohman/helpers"
	"github.com/ainurbrr/task-5-vix-btpns-Moh.AinurBahtiarRohman/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var user models.User
			authHeader := c.Request().Header.Get("Authorization")

			if !strings.Contains(authHeader, "Bearer") {
				response := helpers.ApiResponse(http.StatusUnauthorized, "error", nil, "Unauthorized")
				return c.JSON(http.StatusUnauthorized, response)
			}

			// Split the token and take the value
			tokenString := ""
			dataToken := strings.Split(authHeader, " ")
			if len(dataToken) == 2 {
				tokenString = dataToken[1]
			}

			token, err := helpers.ValidateToken(tokenString)
			if err != nil {
				response := helpers.ApiResponse(http.StatusUnauthorized, "error", nil, "Unauthorized")
				return c.JSON(http.StatusUnauthorized, response)
			}

			claim, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				response := helpers.ApiResponse(http.StatusUnauthorized, "error", nil, "Unauthorized")
				return c.JSON(http.StatusUnauthorized, response)
			}

			userID := int(claim["user_id"].(float64))

			err = db.First(&user, userID).Error
			if err != nil {
				response := helpers.ApiResponse(http.StatusUnauthorized, "error", nil, "Unauthorized")
				return c.JSON(http.StatusUnauthorized, response)
			}

			c.Set("currentUser", user)
			return next(c)
		}
	}
}

