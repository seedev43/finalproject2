package main

import (
	"fp2/database"
	"fp2/routers"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")

	_, err := database.InitDB()

	if err != nil {
		panic(err)
	}

	routers.SetupRouter().Run(PORT)

}
