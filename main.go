package main

import (
	"os"

	"github.com/saschagrunert/crio-demos/pkg/runs"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})

	app := cli.NewApp()
	app.Name = "crio-demos"
	app.Usage = "CRI-O Demonstration Examples"
	app.Authors = []cli.Author{
		{Name: "Sascha Grunert", Email: "sgrunert@suse.com"},
	}
	app.HideVersion = true
	app.UseShortOptionHandling = true
	app.Commands = []cli.Command{
		{Name: "1-interaction", Aliases: []string{"1"}, Action: runs.Interaction},
		{Name: "2-logging", Aliases: []string{"2"}, Action: runs.Logging},
		{Name: "3-lifecycle", Aliases: []string{"3"}, Action: runs.LifeCycle},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
