package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"fmt"
)

func valuesfromtheEnv() {
	app := cli.NewApp()

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "lang, l",
			Value: "english",
			Usage: "language for the greeting",
			//The EnvVar may also be given as a comma-delimited "cascade", where the first environment variable that resolves is used as the default.
			EnvVar: "LEGACY_COMPAT_LANG,APP_LANG,LANG",
		},
	}

	app.Action = func(c *cli.Context) error {
		name := "Nefertiti"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		if c.String("lang") == "code" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}