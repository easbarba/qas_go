package actions

import (
	"fmt"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
)

// TODO: After grabbing log informations

func Grab() {
	fmt.Printf("Grabbing at:  %s \n", projectsDir())

	_, err := git.PlainClone("/tmp/go-git", false, &git.CloneOptions{
		URL:      "https://github.com/go-git/go-git",
		Progress: os.Stdout,
	})

	CheckIfError(err)
}

func CheckIfError(err error) {
	fmt.Println(err)
}

func projectsDir() string {
	home, _ := os.UserHomeDir()

	result := path.Join(home, "Projects")

	return result
}
