package main

import (
	"fmt"
	"os"

	"github.com/caius/hcp-cli/internal/hcp-packer"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("hcp-packer", "0.0.1")
	c.Autocomplete = true
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"list-buckets": func() (cli.Command, error) {
			return &hcp_packer.ListBucketCommand{}, nil
		},
		"version": func() (cli.Command, error) {
			return &hcp_packer.VersionCommand{}, nil
		},
	}

	exitCode, err := c.Run()
	if err != nil {
		fmt.Println("Error executing cli: ", err)
		os.Exit(1)
	}

	os.Exit(exitCode)
}
