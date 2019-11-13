# CRI-O Demos

This repository aims to provide you the demo material for the Kubernetes
container runtime [CRI-O][0].

[0]: https://github.com/cri-o/cri-o

## Requirements

The following requirements needs to be fulfilled to run the demos inside this
repository:

- A running Kubernetes cluster (via systemd units) with CRI-O as configured
  container runtime
  - CRI-O `log_level = "debug"`
- A working Kubeconfig to access that cluster in `$HOME/.kube/config`
- A working `crictl` binary and [configuration][1]

[1]: https://github.com/kubernetes-sigs/cri-tools/blob/master/docs/crictl.md

## How it works

`crio-demos` is a golang based command line application which contains
subcommands for every demo. This make the interactive usage possible whereas
every step has to be confirmed via `ENTER`.

![](.github/demo.svg)

A simple demo may look like this in source code:

```go
package some_package

import (
	"github.com/saschagrunert/crio-demos/pkg/demo"
	"github.com/urfave/cli"
)

func Demo(context *cli.Context) {
	d := demo.New("My demo")

	d.Step([]string{
		"Possible multi-line",
		"description text of the step",
	},
		// Command of the step
		"echo Hello World",
	)

	d.Run()
}
```

## Contents

TBD
