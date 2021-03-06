package main

import (
	"github.com/codegangsta/cli"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	newApp().Run(os.Args)
}

func newApp() *cli.App {

	app := cli.NewApp()
	app.Name = "hn"
	app.Version = Version
	app.Authors = []cli.Author{{Name: "qube81"}}
	app.Usage = "A Hacker News CLI reader"
	app.Commands = Commands

	return app
}
