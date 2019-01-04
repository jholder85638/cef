package cmd

import (
	"fmt"
	"os"

	"github.com/jholder85638/toolbox/atexit"
)

func createDir(dir string, mode os.FileMode) {
	if err := os.MkdirAll(dir, mode); err != nil {
		fmt.Println(err)
		fmt.Println("You may need to run the 'cef' tool as root.")
		atexit.Exit(1)
	}
}

func checkFileError(err error, op, name string) {
	if err != nil {
		fmt.Printf("Unable to %s file %s\n", op, name)
		fmt.Println(err)
		atexit.Exit(1)
	}
}
