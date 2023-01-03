package main

import (
	"github.com/crispybaccoon/hayashi/cli"
	"os"
)

func main() {

	config := cli.Read()

	err := cli.Start(config)
	cli.Err(err)

	os.Exit(0)
}
