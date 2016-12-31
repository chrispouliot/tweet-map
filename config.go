package main

import (
	"fmt"
	"os"
)

var (
	consumerKey    string
	consumerSecret string
	token          string
	tokenSecret    string
)

var logger *Logger

func loadConfig(val *string, envName string, required bool) {
	*val = os.Getenv(envName)
	if required && *val == "" {
		panic(fmt.Errorf("%s not defined", envName))
	}
}

func initConfig() {
	// Required
	loadConfig(&consumerKey, "CONSUMER_KEY", true)
	loadConfig(&consumerSecret, "CONSUMER_SECRET", true)
	loadConfig(&token, "TOKEN", true)
	loadConfig(&tokenSecret, "TOKEN_SECRET", true)

	logger = NewLogger()
}
