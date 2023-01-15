package config

import (
	"log"
	"os"
)

// Home folder of user
func Home() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return home
}
