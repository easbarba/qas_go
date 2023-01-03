package actions

import (
	"fmt"
	"strings"
)

func New(newConfig *string, verbose *bool) {
	newOne := strings.Split(*newConfig, ",")

	fmt.Println(newOne)
}
