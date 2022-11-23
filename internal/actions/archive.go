package actions

import (
	"fmt"
	"path"
	"strings"

	"github.com/easbarba/qas/internal/config"
)

// Archive will zip repositories and place $DOWNLOADS/archived
func Archive(rawlist *string) {
	down := path.Join(config.Home(), "Downloads", "archived")
	list := strings.Split(*rawlist, ",")
	projects := config.All()

	fmt.Println()
	fmt.Printf("Archiving at %s", down)
	fmt.Println()

	for _, project := range projects {
		for _, p := range project.Projects {
			for _, m := range list {
				if p.Name == path.Base(m) {
					fmt.Println()
					fmt.Println(m)
				}
			}

		}
	}
}
