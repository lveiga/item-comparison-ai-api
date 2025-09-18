package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	BindAddr     string
	DatabasePath string
	Environment  string
}

// New - responsible to store env configs
func New() *AppConfig {
	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Println("WARN - ERROR TO LOAD .ENV FILE")
	}

	return &AppConfig{
		BindAddr:     os.Getenv("BIND_ADDR"),
		DatabasePath: os.Getenv("DATA_FILE_PATH"),
		Environment:  os.Getenv("ENVIRONMENT"),
	}
}
