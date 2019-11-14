package setup

import (
	. "github.com/saschagrunert/crio-demos/pkg/demo"
	"github.com/urfave/cli"
)

func EnsureInfoLogLevel() {
	Ensure(
		S(`sudo sed -i -E 's/(log_level = )(.*)/\1"info"/' /etc/crio/crio.conf`),
		S("sudo kill -HUP $(pgrep crio)"),
	)
}

func Before(ctx *cli.Context) error {
	Ensure(
		// Set log_level to debug
		S(`sudo sed -i -E 's/(log_level = )(.*)/\1"debug"/' /etc/crio/crio.conf`),
		S("sudo kill -HUP $(pgrep crio)"),

		// Remove all events
		S("kubectl delete events --all"),
	)
	cleanup()
	return nil
}

func After(ctx *cli.Context) error {
	cleanup()
	return nil
}

func cleanup() {
	Ensure(
		S("sudo pkill kubectl"),
		S("kubectl delete pod nginx alpine"),
		S("kubectl delete deploy nginx"),
	)
}
