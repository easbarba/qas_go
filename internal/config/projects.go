package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"time"
)

// Structure structure of Configuration files
// log config files found
type Structure []struct {
	Lang     string `json:"lang"`
	Projects []struct {
		Name   string `json:"name"`
		Branch string `json:"branch"`
		URL    string `json:"url"`
	} `json:"projects"`
}

// TODO: get from url from ENV Variable or cli
const url = "http://localhost:5000/v1/config/all"

// HomeFolder that all projects repositories will be stored at
var HomeFolder string = path.Join(Home(), "Projects")

// All configuration consumed from the qas API
func All(verbose *bool) Structure {
	qasClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "qas")

	res, getErr := qasClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	config := Structure{}
	jsonErr := json.Unmarshal(body, &config)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return config
}
