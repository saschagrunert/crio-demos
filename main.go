package main

import (
	"os"

	"github.com/saschagrunert/crio-demos/pkg/runs"
	"github.com/saschagrunert/crio-demos/pkg/setup"
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
	app.Before = setup.Before
	app.After = setup.After
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "auto",
			Usage: "run the demo in automatic mode",
		},
	}
	app.Commands = []cli.Command{
		{Name: "1-interaction", Aliases: []string{"1"}, Action: runs.Interaction},
		{Name: "2-logging", Aliases: []string{"2"}, Action: runs.Logging},
		{Name: "3-lifecycle", Aliases: []string{"3"}, Action: runs.LifeCycle},
		{Name: "4-port-forward", Aliases: []string{"4"}, Action: runs.PortForward},
		{Name: "5-recovering", Aliases: []string{"5"}, Action: runs.Recovering},
		{Name: "6-networking", Aliases: []string{"6"}, Action: runs.Networking},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
