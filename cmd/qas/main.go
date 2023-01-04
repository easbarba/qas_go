package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/easbarba/qas/internal/actions"
)

var verbose = flag.Bool("verbose", false, "display more information")
var grab = flag.Bool("grab", false, "grab floss projects")
var archive = flag.String("archive", "", "archive floss projects listed on NAMES")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "qas - Easily manage multiple FLOSS repositories. \n")
		fmt.Fprintln(flag.CommandLine.Output(), "\nUsage information:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if !*grab && *archive == "" {
		flag.Usage()
		os.Exit(0)
	}

	if *grab {
		actions.Grab(verbose)
		return
	}

	actions.Archive(archive, verbose)
}
