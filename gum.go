package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/xLegoz/gum/cmds"

	_ "github.com/xLegoz/gum/services/datastores"
	_ "github.com/xLegoz/gum/services/languages"
)

func main() {
	app := cli.NewApp()
	app.Name = "gum"
	app.Usage = "manage your application environments"

	app.Commands = []cli.Command{
		{
			Name:     "init",
			Usage:    "create a gum application",
			Category: "maintenance",
			Action:   cmds.MaintenanceInit,
		},
		{
			Name:     "clean",
			Usage:    "clean out all gum setup for this repository (leaves configurations)",
			Category: "maintenance",
			Action:   cmds.MaintenanceClean,
		},

		{
			Name:     "status",
			Usage:    "status of the gum application",
			Category: "operations",
			Action:   cmds.OperationStatus,
		},
		{
			Name:     "up",
			Usage:    "starts the gum application",
			Category: "operations",
			Action:   cmds.OperationUp,
		},
		{
			Name:     "restart",
			Usage:    "restarts the gum application",
			Category: "operations",
			Action:   cmds.OperationRestart,
		},
		{
			Name:     "down",
			Usage:    "brings the gum application down",
			Category: "operations",
			Action:   cmds.OperationDown,
		},
		{
			Name:     "logs",
			Usage:    "show the logs of the gum application",
			Category: "operations",
			Action:   cmds.OperationLogs,
		},
	}
	app.Run(os.Args)
}
