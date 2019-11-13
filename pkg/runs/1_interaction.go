package runs

import (
	. "github.com/saschagrunert/crio-demos/pkg/demo"
	"github.com/urfave/cli"
)

func Interaction(context *cli.Context) {
	d := New("Demo 1 - Basic interactions with CRI-O")

	d.Step(X(
		"The recommended way of running CRI-O is as a systemd unit.",
		"Let's verify that CRI-O is running as expected",
	), X(
		"sudo systemctl --no-pager status crio",
	))

	d.Step(X(
		"If CRI-O is up and running, then a kubelet instance can",
		"be configured to run CRI-O",
	), X(
		"sudo systemctl --no-pager status kubelet",
	))

	d.Step(X(
		"We should be now able to interact with CRI-O via `crictl`",
	), X(
		"sudo crictl version",
	))

	d.Step(X(
		"We can list the pods and their status",
	), X(
		"sudo crictl pods",
	))

	d.Step(X(
		"Or the containers",
	), X(
		"sudo crictl ps -a",
	))

	d.Step(X(
		"All crictl calls result in direct gRPC request to CRI-O",
		"For example, `crictl ps` results in a `ListContainersRequest`.",
	), X(
		"sudo journalctl -u crio --since '1 minute ago' |",
		"grep -Po '.*ListContainers(Request|Response){.*?}'",
	))

	d.Step(X(
		"It looks like that the kubelet syncs periodically with CRI-O.",
		"Let's check that",
	), X(
		"sudo journalctl -fu crio",
	))

	d.Run()
}
