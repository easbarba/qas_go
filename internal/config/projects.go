package config

import (
	// "encoding/csv"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// log config files found

func Find() {
	files, err := ioutil.ReadDir(ConfigFolder())

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		p := path.Join(ConfigFolder(), file.Name())

		if _, err := os.Stat(p); err != nil {
			// fmt.Println(err.Error())
			continue
		}

		parse(p)
	}
}

func parse(filepath string) {
	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	fmt.Println(records)
}

func home() string {
	home, _ := os.UserHomeDir()

	return home
}

func Folder() string {
	result := path.Join(home(), "Projects")

	return result
}

func ConfigFolder() string {
	result := path.Join(home(), ".config", "qas")

	return result
}
