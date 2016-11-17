package cmds

import (
  "fmt"

  "github.com/urfave/cli"
)

func MaintenanceInit(c *cli.Context) error {
  fmt.Println("added task: ", c.Args().First())
  return nil
}

func MaintenanceClean(c *cli.Context) error {
  fmt.Println("added task: ", c.Args().First())
  return nil
}
