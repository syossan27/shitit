package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := makeApp()
	app.Run(os.Args)
}

func makeApp() *cli.App {
	app := cli.NewApp()
	app.Name = "shitit"
	app.Usage = "💩"
	app.Version = "9999"
	app.Action = HolyShit

	return app
}
