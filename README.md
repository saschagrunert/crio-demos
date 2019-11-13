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

## Contents

1. [Basic interactions with CRI-O](pkg/runs/1_interaction.go)
2. [Logging and live configuration reload](pkg/runs/2_logging_live_reload.go)

## How it works

`crio-demos` is a golang based command line application which contains
subcommands for every demo. This make the interactive usage possible whereas
every step has to be confirmed via `ENTER`.

![](.github/demo.svg)

A simple demo may look like this in source code:

```go
package some_package

import (
    . "github.com/saschagrunert/crio-demos/pkg/demo"
    "github.com/urfave/cli"
)

func Demo(context *cli.Context) {
    // Preparation steps won't be printed, they're just there
    // to setup a pre-defined environment
    Prepare(
        S("echo Preparing..."),
        S("echo Hello"),
    )

    // A new demo has a title printed at the beginning of the run
    d := New(
        "My Demo",
        "This is just a demo",
    )

    // A demo can consists of multiple steps, each step has a description and
    // a command to be executed.
    d.Step(S(
        "This is the description of the step,",
        "which supports multiple lines, too",
    ), S(
        "echo Hello world",
    ))

    d.Run()
}
```
