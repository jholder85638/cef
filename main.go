package main

import (
	"fmt"
	"os"

	"github.com/jholder85638/cef/internal/cmd"
	"github.com/jholder85638/toolbox/atexit"
	"github.com/jholder85638/toolbox/cmdline"
)

const desiredCEFVersion = "3.3538.1852.gcb937fc"

func main() {
	cmdline.CopyrightYears = "2018"
	cmdline.CopyrightHolder = "Richard A. Wilkes"
	cmdline.AppIdentifier = "com.trollworks.cef"
	cl := cmdline.New(true)
	cl.Description = "Utilities for managing setup of the Chromium Embedded Framework."
	cl.AddCommand(cmd.NewInstall(desiredCEFVersion))
	cl.AddCommand(cmd.NewDist())
	if err := cl.RunCommand(cl.Parse(os.Args[1:])); err != nil {
		fmt.Fprintln(os.Stderr, err)
		atexit.Exit(1)
	}
}
