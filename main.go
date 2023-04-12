package main

import (
	"challenge-2/database"
	"challenge-2/router"
	"os"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	var PORT = os.Getenv("PORT")
	// r.Run(":8081")
	r.Run(":" + PORT)

}
