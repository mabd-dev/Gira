package main

import (
	"fmt"
	"os"

	"github.com/mabd-dev/gira/api"
	"github.com/mabd-dev/gira/config"
	"github.com/mabd-dev/gira/internal/logger"
	"github.com/mabd-dev/gira/internal/ui"
)

const version = "0.1.1"

func main() {
	// Check for version command
	if len(os.Args) > 1 {
		arg := os.Args[1]
		if arg == "version" || arg == "-v" || arg == "--version" {
			fmt.Printf("gira version %s\n", version)
			os.Exit(0)
		}
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configData, err := config.Load(homeDir)
	if err != nil {
		if myError, ok := err.(*config.FirstTimeError); ok {
			fmt.Println(myError.Message)
			os.Exit(0)
		}
		panic(err)
	}

	if configData.General.Debug {
		if err := createMockAPIClient(); err != nil {
			panic(err)
		}
	} else {
		if err := createRealApiClient(configData.Credentials); err != nil {
			panic(err)
		}
	}

	logger.Init(true, "/.config/gira/logs/")

	if err := ui.Render(); err != nil {
		// fmt.Printf("failed to render using bubbletea, err=%s", err.Error())
		panic(err)
	}
}

func createMockAPIClient() error {
	_, err := api.NewMockClient("")
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func createRealApiClient(cred config.CredentialsConfig) error {
	_, err := api.NewClient(
		cred.Email,
		cred.Secret,
		cred.Domain,
	)
	if err != nil {
		fmt.Printf("error creating new client, err=%s", err.Error())
		return err
	}
	return nil
}
