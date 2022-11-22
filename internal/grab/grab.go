package grab

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

// TODO: After grabbing log informations

func Grab() {
	fmt.Println("grabbing")

	_, err := git.PlainClone("/tmp/go-git", false, &git.CloneOptions{
		URL:      "https://github.com/go-git/go-git",
		Progress: os.Stdout,
	})
	CheckIfError(err)
}

func CheckIfError(err error) {
	fmt.Println(err)
}
