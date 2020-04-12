package main

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/noukenolife/authserver/di"
	"github.com/noukenolife/authserver/helper"
)

func main() {
	godotenv.Load()

	serverAddress := helper.GetEnvWithDefaultValue("SERVER_ADDRESS", "https://localhost")
	parsedServerAddress, err := url.Parse(serverAddress)
	if err != nil {
		panic(serverAddress + " is not a valid address.")
	}

	r := gin.Default()

	container, err := di.NewContainer()
	if err != nil {
		panic("Failed to create a container")
	}

	container.Router.InitRoutes(r)

	fmt.Println("Running: " + parsedServerAddress.String() + "\n")
	r.Run(parsedServerAddress.Host)
}
