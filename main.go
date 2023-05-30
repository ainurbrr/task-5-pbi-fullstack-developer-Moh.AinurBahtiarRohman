package main

import (
	"github.com/ainurbrr/task-5-vix-btpns-Moh.AinurBahtiarRohman/database"
	"github.com/ainurbrr/task-5-vix-btpns-Moh.AinurBahtiarRohman/router"
)

func main() {
	database.Init()
	e := router.RouteInit()
	e.Start(":8080")
}
