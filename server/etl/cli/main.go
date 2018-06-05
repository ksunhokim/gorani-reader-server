package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	addrFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "addr",
			Value: "localhost:5982",
			Usage: "address of the server",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:      "dict",
			ArgsUsage: "[dict path]",
			Flags:     addrFlags,
			Usage:     "post words to etl server from dict json file",
			Action: func(c *cli.Context) error {
				dict := c.Args().First()
				addr := c.String("addr")
				return dictToServer(addr, dict)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
