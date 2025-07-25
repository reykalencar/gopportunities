package main

import (
	"github.com/reykalencar/gopportunities/config"
	"github.com/reykalencar/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	//Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initializtion error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
