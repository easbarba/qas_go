package actions

import (
	"fmt"
	"path"
	"strings"

	"github.com/easbarba/qas/internal/config"
)

var archiveFolder string = path.Join(config.Home(), "Downloads", "archived")

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
	fmt.Println("Archiving:", project)
}
