package runs

import (
	. "github.com/saschagrunert/crio-demos/pkg/demo"
	"github.com/urfave/cli"
)

func Recovering(context *cli.Context) {
	d := New("Recovering Workloads")

	d.Step(S(
		"Let’s start with a fresh nginx deployment",
	), S(
		"kubectl create deployment --image=nginx nginx",
	))

	d.Step(S(
		"Now we kill the container internal nginx process",
	), S(
		"sudo pkill -KILL nginx",
	))

	d.Step(S(
		"Then the container monitor conmon will notice that and the workload gets removed",
	), S(
		"sudo journalctl -u crio --since '1 minute ago' | grep -A1 exited",
	))

	d.Step(S(
		"The kubelets synchronization loop will notice that as well and will re-schedule the workload",
	), S(
		"sudo journalctl -u kubelet --since '1 minute ago' | grep -A1 ContainerDied",
	))

	d.Run()
}
