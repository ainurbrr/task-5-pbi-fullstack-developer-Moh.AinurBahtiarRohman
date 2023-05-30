package main

import "task-5-vix-btpns-Moh.AinurBahtiarRohman/router"
import "task-5-vix-btpns-Moh.AinurBahtiarRohman/database"

func main() {
	database.Init()
	r := router.RouteInit()
	r.Run(":8080")
}
