package runs

import (
	. "github.com/saschagrunert/crio-demos/pkg/demo"
	"github.com/urfave/cli"
)

func PortForward(context *cli.Context) {
	Run(
		S(`sudo sed -i -E 's/(log_level = )(.*)/\1"debug"/' /etc/crio/crio.conf`),
		S("sudo kill -HUP $(pgrep crio)"),
	)

	d := New("Port Forwarding")

	d.Step(S(
		"First, let’s create a workload which we want to access",
	), S(
		"kubectl run --generator=run-pod/v1 --image=nginx nginx",
	))

	d.Step(S(
		"Then, a port-forward can be done using kubectl",
	), S(
		"kubectl port-forward pod/nginx 8888:80 >/dev/null &",
	))

	d.Step(S(
		"Now we're able to access the pods nginx server",
	), S(
		"curl 127.0.0.1:8888",
	))

	d.Step(S(
		"During port forward, CRI-O returns a streaming endpoint to the kubelet",
	), S(
		"sudo journalctl -u crio --since '2 minutes ago' | grep -E '(PortForward(Request|Response)|socat).*'",
	))

	d.Step(S(
		"So we cou use `socat` directly to access the web server after entering the PID namespace",
	), S(
		`echo "GET /" |`,
		` sudo $(sudo journalctl -u crio --since '2 minute ago' |`,
		`sed -n -E 's;.*executing port forwarding command: (.*80).*;\1;p')`,
	))

	d.Run()

	Run(
		S("sudo pkill kubectl"),
		S("kubectl delete pod nginx"),
	)
}