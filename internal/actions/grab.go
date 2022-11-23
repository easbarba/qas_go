package actions

import (
	"fmt"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"

	"github.com/easbarba/qas/internal/config"
)

// TODO: After grabbing informations log

// Grab all project by pulling or cloning
func Grab() {
	projects := config.All()

	for _, project := range projects {
		for _, p := range project.Projects {
			fld := path.Join(config.HomeFolder(), project.Lang, p.Name)

			printInfo(fld, p.Name, p.URL, p.Branch)

			if _, err := os.Stat(path.Join(fld, ".git")); err == nil {
				pull(fld, p.URL, p.Branch)
			} else {
				clone(fld, p.Name, p.URL, p.Branch)
			}

		}
	}

	// TODO return error
}

func printInfo(folder, name, url, branch string) {
	in := `
name: %s
url: %s
branch: %s
folder: %s
`
	fmt.Printf(in, name, url, branch, folder)
}

func clone(folder, name, url, branch string) {
	fmt.Println("status: cloning")
	fmt.Println("")
	branch = fmt.Sprintf("refs/heads/%s", branch)

	_, err := git.PlainClone(folder, false, &git.CloneOptions{
		URL:           url,
		ReferenceName: plumbing.ReferenceName(branch),
		Progress:      os.Stdout,
		SingleBranch:  true,
		Depth:         1,
	})

	CheckIfError(err)
}

func pull(folder, url, branch string) {
	fmt.Println("status: pulling")
	fmt.Println("")

	o, err := git.PlainOpen(folder)
	CheckIfError(err)

	w, err := o.Worktree()
	CheckIfError(err)

	w.Pull(&git.PullOptions{
		RemoteName:    "origin",
		ReferenceName: plumbing.ReferenceName(branch),
		SingleBranch:  true,
		Depth:         1,
		Progress:      os.Stdout,
	})
	CheckIfError(err)
}
