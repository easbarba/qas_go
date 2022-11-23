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

// log config files found
type Config struct {
	Lang     string `json:"lang"`
	Projects []struct {
		Name   string `json:"name"`
		Branch string `json:"branch"`
		URL    string `json:"url"`
	} `json:"projects"`
}

func All() []Config {
	var result []Config
	files := files()

	for _, file := range files {
		p := path.Join(Folder(), file.Name())

		// symbolic link is broken
		if _, err := os.Stat(p); err != nil {
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

// array of configuratiion file name
func files() []fs.FileInfo {
	files, err := ioutil.ReadDir(Folder())

	if err != nil {
		log.Fatal(err)
	}

	return files
}

func parse(filepath string) Config {
	file, err := ioutil.ReadFile(filepath)
	var proj Config

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(file, &proj)

	if err != nil {
		fmt.Println(err)
	}

	return proj
}

func home() string {
	home, _ := os.UserHomeDir()

	return home
}

// folder that all projects repositories will be stored at
func HomeFolder() string {
	result := path.Join(home(), "Projects")

	return result
}

// configuration folder that config files will be looked up for
func Folder() string {
	result := path.Join(home(), ".config", "qas")

	return result
}
