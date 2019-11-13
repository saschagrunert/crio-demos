package runs

import (
	. "github.com/saschagrunert/crio-demos/pkg/demo"
	"github.com/urfave/cli"
)

func Logging(context *cli.Context) {
	Prepare(
		S(`sudo sed -i -E 's/(log_level = )(.*)/\1"info"/' /etc/crio/crio.conf`),
		S("sudo kill -HUP $(pgrep crio)"),
	)

	d := New(
		"Demo 2 - Logging and configuration reload",
		"This demo shows how to configure CRI-O logging and",
		"reload the configuration during runtime",
	)

	d.Step(S(
		"The basic configuration file of CRI-O is available in",
		"/etc/crio/crio.conf",
	), S(
		"less /etc/crio/crio.conf",
	))

	d.Step(S(
		"For example, the log level can be changed there too",
	), S(
		"grep -B3 log_level /etc/crio/crio.conf",
	))

	d.Step(S(
		"So we can set the `log_level` to a higher verbosity",
	), S(
		`sudo sed -i -E 's/(log_level = )(.*)/\1"debug"/' /etc/crio/crio.conf`,
	))

	d.Step(S(
		"To reload CRI-O, we have to send a SIGHUP (hangup) to the process",
	), S(
		"sudo kill -HUP $(pgrep crio)",
	))

	d.Step(S(
		"The logs indicate that the configuration has been reloaded correctly",
	), S(
		"sudo journalctl -u crio --since '1 minute ago' |",
		"grep -B1 'log_level.*debug'",
	))

	d.Step(S(
		"CRI-O now logs every request and response as seen in the first demo",
	), S(
		"sudo journalctl -fu crio",
	))

	d.Run()
}
