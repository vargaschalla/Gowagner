package main

import (
	"os"

	"github.com/vargaschalla/Gowagner/routers"
)

func main() {

	r := routers.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	r.Run(":" + port) //"localhost:8081"
}
