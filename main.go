package main

import (
	"os"

	"github.com/tim-lynn-clark/tlccryptotool/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
