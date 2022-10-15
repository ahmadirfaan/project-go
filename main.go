package main

import (
	"github.com/ahmadirfaan/project-go/app"
	"github.com/ahmadirfaan/project-go/cli"
	"os"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())
}
