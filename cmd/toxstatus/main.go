package main

import (
	"os"

	"github.com/2mf/ToxStatus/cmd/toxstatus/cmd"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		os.Exit(1)
	}
}
