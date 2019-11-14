package main

import (
	"os"
	"os/signal"

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
			Name: "auto, a",
			Usage: "run the demo in automatic mode, " +
				"where every step gets executed automatically after 5 seconds",
		},
		cli.BoolFlag{
			Name:  "continuously, c",
			Usage: "run the demo continuously",
		},
	}
	app.Commands = []cli.Command{
		{Name: "all", Aliases: []string{"a"}, Action: func(ctx *cli.Context) {
			run := func() {
				runs.Interaction(ctx)
				runs.Logging(ctx)
				runs.LifeCycle(ctx)
				runs.PortForward(ctx)
				runs.Recovering(ctx)
				runs.Networking(ctx)
			}
			if ctx.GlobalBool("continuously") {
				for {
					run()
				}
			} else {
				run()
			}
		}},
		{Name: "1-interaction", Aliases: []string{"1"}, Action: runs.Interaction},
		{Name: "2-logging", Aliases: []string{"2"}, Action: runs.Logging},
		{Name: "3-lifecycle", Aliases: []string{"3"}, Action: runs.LifeCycle},
		{Name: "4-port-forward", Aliases: []string{"4"}, Action: runs.PortForward},
		{Name: "5-recovering", Aliases: []string{"5"}, Action: runs.Recovering},
		{Name: "6-networking", Aliases: []string{"6"}, Action: runs.Networking},
		{Name: "7-pull-auth", Aliases: []string{"7"}, Action: runs.PullAuth},
	}

	// Catch interrupts
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			_ = setup.After(nil)
			os.Exit(0)
		}
	}()

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
