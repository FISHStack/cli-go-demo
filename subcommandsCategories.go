package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func subcommandsCategories() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name: "noop",
		},
		{
			Name:     "add",
			Category: "Template actions",
		},
		{
			Name:     "remove",
			Category: "Template actions",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
