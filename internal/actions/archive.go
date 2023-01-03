package actions

import (
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/easbarba/qas/internal/config"
)

var archiveFolder string = path.Join(config.Home(), "Downloads", "archived")
var spin = spinner.New(spinner.CharSets[26], 100*time.Millisecond)

// Archive will zip repositories and place $DOWNLOADS/archived
func Archive(rawlist *string, verbose *bool) {
	projects := config.All(verbose)
	list := strings.Split(*rawlist, ",")

	if *verbose {
		fmt.Printf("\nArchiving at %s\n", archiveFolder)
	}

	for _, project := range projects {
		for _, p := range project.Projects {
			for _, m := range list {
				if p.Name == path.Base(m) {
					do(m)
				}
			}
		}
	}
}

func do(project string) {
	spin.Start()
	fmt.Println(project)
	spin.Stop()
}
