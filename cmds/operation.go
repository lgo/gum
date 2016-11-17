package cmds

import (
	"github.com/urfave/cli"
	"github.com/xLegoz/gum/configuration"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func OperationUp(c *cli.Context) error {
	var config = configuration.Configuration{}
	err := config.LoadConfiguration(c.Args().First())
	check(err)
	return nil
}

func OperationStatus(c *cli.Context) error {
	return nil
}

func OperationRestart(c *cli.Context) error {
	return nil
}

func OperationDown(c *cli.Context) error {
	return nil
}

func OperationLogs(c *cli.Context) error {
	return nil
}
