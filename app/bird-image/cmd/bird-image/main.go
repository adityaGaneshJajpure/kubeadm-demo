package main

import (
	"log"
	"os"

	_ "net/http/pprof"

	"get-bird-image/pkg/constants"
	"get-bird-image/pkg/http/httpclient"
	srv "get-bird-image/pkg/http/server"
	"get-bird-image/pkg/logger"
	"get-bird-image/pkg/service"
)

func main() {
	// validate env
	validateEnv()

	loggerObj, err := logger.NewLogger()
	if err != nil {
		log.Panic("error while creating logger: ", err)
	}
	defer loggerObj.DisconnectLogger()

	// init gin server
	server := srv.Server{Logger: loggerObj}
	serverInitErr := server.InitServer()
	if serverInitErr != nil {
		log.Panic("error initializing server:", serverInitErr)
	}

	serviceHandler := service.NewServiceHandler(loggerObj, &httpclient.HTTPClient{})
	serverRouteErr := server.SetupRoutes(serviceHandler)
	if serverRouteErr != nil {
		log.Panic("error setting up routes in server:", serverRouteErr)
	}
	// run server
	if err := server.RunServer(); err != nil {
		log.Panic("unable to run the server!")
	}
}

func validateEnv() {
	envs := []string{
		constants.PORT,
		constants.UNSPLASH_API,
		constants.UNSPLASH_API_KEY,
		constants.DEFAULT_BIRD_IMAGE,
	}
	for _, env := range envs {
		if os.Getenv(env) == "" {
			log.Panicf("Env %s is missing", env)
		}
	}
}
