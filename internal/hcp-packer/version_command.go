package hcp_packer

import (
	"fmt"
)

type VersionCommand struct{}

func (c *VersionCommand) Run(args []string) int {
	fmt.Println("hcp-packer v0.0.1")
	return 0
}

func (c *VersionCommand) Help() string {
	return "Show the hcp-packer version"
}

func (c *VersionCommand) Synopsis() string {
	return c.Help()
}
