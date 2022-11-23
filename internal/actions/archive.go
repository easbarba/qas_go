package actions

import (
	"fmt"
	"strings"
)

// Archive will zip repositories and place $DOWNLOADS/archive
func Archive(projects *string) {
	x := strings.Split(*projects, ",")
	fmt.Printf("archiving %s", x)
}
