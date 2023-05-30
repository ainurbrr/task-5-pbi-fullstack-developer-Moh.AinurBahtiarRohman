package router

import (
	"task-5-vix-btpns-Moh.AinurBahtiarRohman/controllers"
	"task-5-vix-btpns-Moh.AinurBahtiarRohman/middlewares"
)

func RouteInit() *echo.Engine {
	route := echo.Default()
	route.Static("/images", "./static/images")

	db := database.GetDB()

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

	photoRoute := api.Group("/photo")
	{
		photoRoute.GET("/", middlewares.AuthMiddleware(db), photoController.Get)
		photoRoute.POST("/", middlewares.AuthMiddleware(db), photoController.Create)
		photoRoute.PUT("/", middlewares.AuthMiddleware(db), photoController.Update)
		photoRoute.DELETE("/", middlewares.AuthMiddleware(db), photoController.Delete)
	}

	return route
}
