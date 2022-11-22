package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/easbarba/qas.go/internal/archive"
	"github.com/easbarba/qas.go/internal/grab"
)

func main() {
	grabValue, archiveValue := parse()

	if *grabValue == true {
		grab.Grab()
	}

	if *archiveValue != "" {
		archive.Archive(archiveValue)
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
