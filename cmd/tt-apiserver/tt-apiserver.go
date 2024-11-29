package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pastry-personal5/tt-apiserver/internal/config"
	"github.com/pastry-personal5/tt-apiserver/internal/models"
	"github.com/pastry-personal5/tt-apiserver/internal/routers"
	"gopkg.in/yaml.v3"
)

func main() {
	applicationName := "tt-apiserver"
	globalConfigFile := "./configs/global_config.yaml"

	file, err := os.Open(globalConfigFile)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var global_config config.GlobalConfig

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&global_config)
	if err != nil {
		log.Fatalf("Error decoding YAML: %v", err)
	}

	fmt.Printf("Starting %v...\n", applicationName)
	config.ConnectDB(global_config)
	// Migrate the User model
	config.DB.AutoMigrate(&models.ExpenseTransaction{})

	r := routers.SetupRouter()

	r.Run(":8080") // Default HTTP server port
}
