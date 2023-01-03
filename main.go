package main

import (
	"github.com/crispybaccoon/hayashi/cli"
	"os"
)

func main() {

	config := cli.Read()

	cli.Start(config)

	os.Exit(0)
}
