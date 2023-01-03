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
func Archive(list *string, verbose *bool) {
	configs := config.All(verbose)
	projectsList := strings.Split(*list, ",")

	if *verbose {
		fmt.Printf("\nArchiving at %s\n", archiveFolder)
	}

	for _, config := range configs {
		for _, p := range config.Projects {
			for _, m := range projectsList {
				if p.Name == path.Base(m) {
					do(m)
				}
			}
		}
	}
}

// TODO: mkdir archive folder
// TODO: archive to zip by default
// TODO: store at archive folder
func do(project string) {
	spin.Start()
	fmt.Println(project)
	spin.Stop()
}
