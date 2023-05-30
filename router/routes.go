package router

import (
	"github.com/ainurbrr/task-5-vix-btpns-Moh.AinurBahtiarRohman/controllers"
	"github.com/ainurbrr/task-5-vix-btpns-Moh.AinurBahtiarRohman/database"
	"github.com/ainurbrr/task-5-vix-btpns-Moh.AinurBahtiarRohman/middlewares"
	"github.com/labstack/echo/v4"
)

func RouteInit() *echo.Echo {
	route := echo.New()
	route.Static("/images", "./static/images")

	db := database.DB

	userController := controllers.NewUserController(db)
	photoController := controllers.NewPhotoController(db)

	api := route.Group("/api/v1")

	userRoute := api.Group("/users")
	{
		userRoute.POST("/register", userController.Register)
		userRoute.POST("/login", userController.Login)
		userRoute.PUT("/:userId", userController.Update)
		userRoute.DELETE("/:userId", userController.Delete)
	}

	photoRoute := api.Group("/photos")
	{
		photoRoute.GET("", photoController.Get, middlewares.AuthMiddleware(db))
		photoRoute.POST("", photoController.Create, middlewares.AuthMiddleware(db))
		photoRoute.PUT("", photoController.Update, middlewares.AuthMiddleware(db))
		photoRoute.DELETE("", photoController.Delete, middlewares.AuthMiddleware(db))
	}

	return route
}
