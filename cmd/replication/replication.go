package main

import (
	"cli/config"
	"log"
	"os"
	"path/filepath"
)


func main() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get the homedir\n")
		return
	}
	
	configPath := filepath.Join(homeDir, ".config", "cli", "config.yml")
	_, err = config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("failed to get the config\n")
		return
	}


	

}