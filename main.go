package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/easbarba/qas/internal/actions"
)

func main() {
	archive, grab, verbose := cliParser()

	if *grab {
		actions.Grab(verbose)
		return
	}

	actions.Archive(archive, verbose)
}

func cliParser() (*string, *bool, *bool) {
	archive := flag.String("archive", "", "archive floss projects listed on NAMES")
	grab := flag.Bool("grab", false, "grab floss projects")
	verbose := flag.Bool("verbose", false, "display more information")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "\nqas - Easily manage multiple FLOSS repositories.\n")
		fmt.Fprintln(flag.CommandLine.Output(), "\nUsage information:")
		flag.PrintDefaults()
	}
	flag.Parse()

	if !*grab && *archive == "" {
		flag.Usage()
		os.Exit(0)
	}

	return archive, grab, verbose
}
