package runs

import (
	. "github.com/saschagrunert/crio-demos/pkg/demo"
	"github.com/urfave/cli"
)

func Interaction(context *cli.Context) {
	Run(
		S(`sudo sed -i -E 's/(log_level = )(.*)/\1"debug"/' /etc/crio/crio.conf`),
		S("sudo kill -HUP $(pgrep crio)"),
	)

	d := New(
		"Demo 1 - Basic interactions with CRI-O",
		"This demo shows basic interactions with CRI-O and",
		"beteen the kubelet and CRI-O.",
	)

	d.Step(S(
		"The recommended way of running CRI-O is as a systemd unit.",
		"Let's verify that CRI-O is running as expected",
	), S(
		"sudo systemctl --no-pager status crio",
	))

	d.Step(S(
		"If CRI-O is up and running, then a kubelet instance can",
		"be configured to run CRI-O",
	), S(
		"sudo systemctl --no-pager status kubelet",
	))

	d.Step(S(
		"We should be now able to interact with CRI-O via `crictl`",
	), S(
		"sudo crictl version",
	))

	d.Step(S(
		"We can list the pods and their status",
	), S(
		"sudo crictl pods",
	))

	d.Step(S(
		"Or the containers",
	), S(
		"sudo crictl ps -a",
	))

	d.Step(S(
		"All crictl calls result in direct gRPC request to CRI-O",
		"For example, `crictl ps` results in a `ListContainersRequest`.",
	), S(
		"sudo journalctl -u crio --since '1 minute ago' |",
		"grep -Po '.*ListContainers(Request|Response){.*?}'",
	))

	d.Step(S(
		"It looks like that the kubelet syncs periodically with CRI-O.",
		"Let's check that",
	), S(
		"sudo journalctl -fu crio",
	))

	d.Run()
}
