package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

// Config structure of Configuration files
// log config files found
type Config struct {
	Lang     string `json:"lang"`
	Projects []struct {
		Name   string `json:"name"`
		Branch string `json:"branch"`
		URL    string `json:"url"`
	} `json:"projects"`
}

// HomeFolder that all projects repositories will be stored at
var HomeFolder string = path.Join(Home(), "Projects")

// folder that config files will be looked up for
var folder string = path.Join(Home(), ".config", "qas")

// All configuration files unmarshallowed
func All(verbose *bool) []Config {
	var result []Config

	files, err := files()
	if err != nil {
		fmt.Println("no configuration file found!")
		os.Exit(1)
	}

	if *verbose {
		// printConfig(files)
	}

	for _, file := range files {
		p := path.Join(folder, file.Name())
		fileInfo, err := os.Stat(p)

		// ignore broken symbolic link
		if os.IsNotExist(err) {
			continue
		}

		// ignore directories
		if fileInfo.IsDir() {
			continue
		}

		// ignore csv files (legacy)
		if ext := filepath.Ext(p); ext == ".csv" {
			continue
		}

		parsed := parse(p)
		result = append(result, parsed)
	}

	return result
}

// Print Configuration minimal
func printConfig(p string) {
	fmt.Println("Configuration files found: ")
	fmt.Print(p)
}

// all configuration files found TODO: return error if no configuration is found.
func files() ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	return files, nil
}

// Parse configuration file, check if the expect syntax is correct TODO: or err.
func parse(filepath string) Config {
	var config Config

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

// Home folder of user
func Home() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return home
}
