package runs

import (
	. "github.com/saschagrunert/crio-demos/pkg/demo"
	"github.com/urfave/cli"
)

func Recovering(context *cli.Context) {
	d := New("Recovering Workloads")

	d.Run()
}
