package cmds

import (
	"os"
	"path/filepath"

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
	dir, err := filepath.Abs("./" + c.Args().First())
	check(err)
	err = os.Chdir(dir)
	check(err)
	err = config.LoadAndCheckConfiguration()
	check(err)
	err = config.Prepare()
	check(err)
	err = config.Start()
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
