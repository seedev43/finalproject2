package main

import (
	"fmt"
	"fp2/database"
	"fp2/routers"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		fmt.Println("port not set. Using default port 8080.")
		PORT = "8080"
	}

	_, err := database.InitDB()

	if err != nil {
		panic(err)
	}

	routers.SetupRouter().Run(":" + PORT)

}
