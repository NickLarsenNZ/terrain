package internal

import (
	"fmt"
	"github.com/nicklarsennz/terrain/utils"
)

type Config struct {
	Diff           bool
	DiffDirectoryA string
	DiffDirectoryB string
	ScanDirectory  string
	ShowResources  bool
	OutputFormat   string
}

var validOutputs = []string{
	"json",
	"yaml",
	"markdown",
	"list",
}

func (c *Config) Validate() []error {
	var errors []error

	utils.AppendError(&errors, c.isOutputValid())

	return errors
}

func (c *Config) isOutputValid() error {
	if !utils.StringInSlice(c.OutputFormat, &validOutputs) {
		return fmt.Errorf("Output Format must be one of %v", validOutputs)
	}
	return nil
}
