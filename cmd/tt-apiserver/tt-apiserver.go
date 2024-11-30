package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pastry-personal5/tt-apiserver/internal/config"
	"github.com/pastry-personal5/tt-apiserver/internal/models"
	"github.com/pastry-personal5/tt-apiserver/internal/routers"
	"github.com/pastry-personal5/tt-apiserver/internal/services"
	"gopkg.in/yaml.v3"
)

func loadConfig(filePath string) (*config.GlobalConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %w", filePath, err)
	}
	defer file.Close()

	var cfg config.GlobalConfig
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("error decoding YAML: %w", err)
	}

	return &cfg, nil
}

func printApplicationStartMessage() {
	const (
		applicationName = "tt-apiserver"
	)
	fmt.Printf("Starting %v...\n", applicationName)
}

func loadGlobalConfigFilePathAndServerPort() (string, string) {
	const (
		defaultGlobalConfigFilePath        = "./configs/global_config.yaml"
		defaultServerPortWithColonPrefixed = ":8080"
	)
	var globalConfigFilePath string

	globalConfigFilePathFromEnv := os.Getenv("GLOBAL_CONFIG_FILE")
	if globalConfigFilePathFromEnv == "" {
		globalConfigFilePath = defaultGlobalConfigFilePath
	} else {
		globalConfigFilePath = globalConfigFilePathFromEnv
	}

	var serverPortWithColonPrefixed string
	portFromEnv := os.Getenv("PORT")
	if portFromEnv == "" {
		serverPortWithColonPrefixed = defaultServerPortWithColonPrefixed
	} else {
		serverPortWithColonPrefixed = ":" + portFromEnv
	}
	return globalConfigFilePath, serverPortWithColonPrefixed
}

func initializeDBConnection(globalConfig *config.GlobalConfig) {

	services.ConnectDB(*globalConfig)
	// Migrate the User model
	services.DB.AutoMigrate(&models.ExpenseTransaction{})
}

func startServerAndRunLoop(serverPortWithColonPrefixed string) {
	r := routers.SetupRouter()
	srv := &http.Server{Addr: serverPortWithColonPrefixed, Handler: r}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}

func main() {
	globalConfigFilePath, serverPortWithColonPrefixed := loadGlobalConfigFilePathAndServerPort()
	globalConfig, err := loadConfig(globalConfigFilePath)
	if err != nil {
		log.Fatalf("Error reading global config file %s: %v", globalConfigFilePath, err)
	}
	printApplicationStartMessage()
	initializeDBConnection(globalConfig)
	startServerAndRunLoop(serverPortWithColonPrefixed)
}
