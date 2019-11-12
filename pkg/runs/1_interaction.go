package runs

import (
	"github.com/saschagrunert/crio-demos/pkg/demo"
	"github.com/urfave/cli"
)

func Interaction(context *cli.Context) {
	d := demo.New("Demo 1 - Interacting with CRI-O")

	d.Step([]string{
		"The recommended way of running CRI-O is as a systemd unit.",
		"Let's verify that CRI-O is running as expected",
	},
		"sudo systemctl status crio",
	)

	d.Step([]string{
		"If CRI-O is up and running, then a kubelet instance can",
		"be configured to run CRI-O",
	},
		"sudo systemctl status kubelet",
	)

	d.Step([]string{
		"We should be now able to interact with CRI-O via `crictl`",
	},
		"sudo crictl version",
	)

	d.Step([]string{
		"We can list the pods and their status",
	},
		"sudo crictl pods",
	)

	d.Step([]string{
		"Or the containers",
	},
		"sudo crictl ps -a",
	)

	d.Run()
}
