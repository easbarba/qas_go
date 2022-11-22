package actions

import (
	"fmt"
	"strings"
)

func Archive(projects *string) {
	x := strings.Split(*projects, ",")
	fmt.Printf("archiving %s", x)
}
