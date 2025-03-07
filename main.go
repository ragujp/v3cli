package main

import (
	"os"

	app "github.com/inonius/v3cli/pkg/client"
)

func main() {
	if err := app.NewCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
