package main

import (
	"dropit/app/globals"
	"dropit/app/http/routes"
	"dropit/configs"
	"dropit/utils"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var Global utils.Map

func init() {
	// application configuration

	// Read the .env file
	content, err := os.ReadFile(".env")
	if err != nil {
		panic(fmt.Sprintf("Couldn't read .env file, due to: %s", err.Error()))
	}

	// Initialize the Environment global struct
	utils.InitEnv(string(content))
	conf := configs.GetConfig()
	//set Log format output
	log.SetFormatter(&log.JSONFormatter{})
	// Tell logrus where to print the output
	log.SetOutput(os.Stdout)
	// minimal log level (report only this level and higher)
	log.SetLevel(conf.LogLevel)
	// Initialize all global variables
	globals.InitGlobals()
}

func main() {
	app := gin.Default()
	// Register all routes
	routes.RegisterRoutes(app)
	// Run server on designated port
	app.Run(fmt.Sprintf(":%s", configs.GetConfig().Port))
}
