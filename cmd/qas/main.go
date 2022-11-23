package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/easbarba/qas/internal/actions"
	"github.com/easbarba/qas/internal/config"
)

func main() {
	grabValue, archiveValue := parse()

	if *grabValue == true {
		actions.Grab(config.All())
	}

	if *archiveValue != "" {
		actions.Archive(archiveValue)
	}
}

// command line arguments parser
func parse() (*bool, *string) {
	grab := flag.Bool("grab", false, "grab floss projects")
	archive := flag.String("archive", "", "archive floss projects listed on NAMES")

	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "qas - Easily manage multiple FLOSS repositories. \n\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	if *grab == false && *archive == "" {
		flag.Usage()
		os.Exit(0)
	}

	return grab, archive
}
