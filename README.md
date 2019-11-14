# CRI-O Demos

This repository aims to provide you the demo material for the Kubernetes
container runtime [CRI-O][0].

[0]: https://github.com/cri-o/cri-o

## Requirements

The following requirements needs to be fulfilled to run the demos inside this
repository:

- A running Kubernetes cluster (via systemd units) with CRI-O as configured
  container runtime
  - CRI-O configuration:
    ```toml
    log_level = "debug"
    cgroup_manager = "systemd"
    conmon_cgroup = "system.slice"
    ```
  - Kubelet should be started with: `-v=2 --cgroup-driver=systemd`
- A working Kubeconfig to access that cluster in `$HOME/.kube/config`
- A working `crictl` binary and [configuration][1]

[1]: https://github.com/kubernetes-sigs/cri-tools/blob/master/docs/crictl.md

## Contents

1. [Basic interactions with CRI-O](pkg/runs/interaction.go)
2. [Logging and live configuration reload](pkg/runs/logging_live_reload.go)
3. [Life Cycle of a Kubernetes workload](pkg/runs/lifecycle.go)
4. [Port Forward](pkg/runs/portforward.go)
5. [Recovering](pkg/runs/recovering.go)
6. [Networking](pkg/runs/networking.go)
7. [Pull Auth](pkg/runs/pull_auth.go)
8. [Registries](pkg/runs/registries.go)
9. [Registry Mirrors](pkg/runs/registry_mirrors.go)
10. [Storage](pkg/runs/storage.go)
