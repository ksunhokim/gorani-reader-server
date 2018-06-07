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

	redisFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "redis",
			Value: "redis://localhost:6379/1",
			Usage: "redis url",
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
		{
			Name:      "relword",
			ArgsUsage: "[reltype]",
			Flags:     addrFlags,
			Usage:     "renew calculated relword",
			Action: func(c *cli.Context) error {
				reltype := c.Args().First()
				addr := c.String("addr")
				return relevantWords(addr, reltype)
			},
		},
		{
			Name:      "addbook",
			ArgsUsage: "[isbn] [epub path]",
			Flags:     append(addrFlags, redisFlags...),
			Usage:     "post book to etl server",
			Action: func(c *cli.Context) error {
				isbn := c.Args().First()
				epub := c.Args().Get(1)
				addr := c.String("addr")
				redisurl := c.String("redis")
				return addBook(isbn, epub, addr, redisurl)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
