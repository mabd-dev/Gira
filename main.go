package main

import (
	"fmt"
	"os"

	"github.com/mabd-dev/gira/api"
	"github.com/mabd-dev/gira/config"
	"github.com/mabd-dev/gira/internal/logger"
	"github.com/mabd-dev/gira/internal/ui"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configData, err := config.Load(homeDir)
	if err != nil {
		if myError, ok := err.(*config.FirstTimeError); ok {
			fmt.Println(myError.Message)
			return
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
	_, err := api.NewMockClient()
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
