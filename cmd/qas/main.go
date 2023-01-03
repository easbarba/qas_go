package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/easbarba/qas/internal/actions"
)

func main() {
	grab, archive, new, verbose := cliParse()

	if *grab == true {
		actions.Grab(verbose)
	}

	if *archive != "" {
		actions.Archive(archive, verbose)
	}

	if *new != "" {
		actions.New(new, verbose)
	}
}

// command line arguments parser
func cliParse() (*bool, *string, *string, *bool) {
	verbose := flag.Bool("verbose", false, "display more information")
	grab := flag.Bool("grab", false, "grab floss projects")
	archive := flag.String("archive", "", "archive floss projects listed on NAMES")
	new := flag.String("new", "", "add a new configuration, eg: js,gum,main,https://github.com/charmbracelet/gum")

	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "qas - Easily manage multiple FLOSS repositories. \n")
		fmt.Fprintln(flag.CommandLine.Output(), "\nUsage information:")
		flag.PrintDefaults()
	}

	if *grab == false && *archive == "" && *new == "" {
		flag.Usage()
		os.Exit(0)
	}

	return grab, archive, new, verbose
}
