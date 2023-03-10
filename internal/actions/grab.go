package actions

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"

	"github.com/easbarba/qas/internal/config"
)

var s = spinner.New(spinner.CharSets[26], 100*time.Millisecond)

// TODO: After grabbing informations log

// Grab all project by pulling or cloning
func Grab(verbose *bool) {
	projects := config.All(verbose)

	for _, project := range projects {
		for _, pj := range project.Projects {
			name := strings.ToLower(pj.Name)
			folder := path.Join(config.HomeFolder, project.Lang, name)

			printInfo(name, pj.URL, pj.Branch, verbose)

			if _, err := os.Stat(path.Join(folder, ".git")); err == nil {
				pull(folder, pj.URL, pj.Branch)
			} else {
				clone(folder, pj.Name, pj.URL, pj.Branch)
			}
		}
	}
	// TODO return error
}

func printInfo(name, url, branch string, verbose *bool) {
	title := color.New(color.FgHiYellow, color.Bold).SprintFunc()
	if *verbose {
		fmt.Print(title("name: "), name, title(" url: "), url, title(" branch: "), branch, "\n")
		return
	}

	fmt.Print(title("name: "), name, "\n")
}

// clone repository if none is found at folder
func clone(folder, name, url, branch string) {
	branch = fmt.Sprintf("refs/heads/%s", branch)

	spin.Start()
	_, err := git.PlainClone(folder, false, &git.CloneOptions{
		URL:           url,
		ReferenceName: plumbing.ReferenceName(branch),
		Progress:      os.Stdout,
		SingleBranch:  true,
		Depth:         1,
	})
	spin.Stop()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// pull repository at url/ and branch in the found folder
func pull(folder, url, branch string) {
	o, err := git.PlainOpen(folder)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	w, err := o.Worktree()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spin.Start()
	w.Pull(&git.PullOptions{
		RemoteName:    "origin",
		ReferenceName: plumbing.ReferenceName(branch),
		SingleBranch:  true,
		Depth:         1,
		Progress:      os.Stdout,
	})
	spin.Stop()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
