package main


import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func valuesfromFiles() {
	app := cli.NewApp()

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "password, p",
			Usage: "password for the mysql database",
			FilePath: "/etc/mysql/password",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}