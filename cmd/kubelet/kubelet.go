// The kubelet binary is responsible for maintaining a set of containers on a particular host VM.
// It syncs data from both configuration file(s) as well as from a quorum of etcd servers.
// It then communicates with the container runtime (or a CRI shim for the runtime) to see what is
// currently running.  It synchronizes the configuration data, with the running set of containers
// by starting or stopping containers.
package main

import (
	"k8s.io/component-base/cli"
	"k8s.io/kubernetes/cmd/kubelet/app"
	"os"
)

func main() {
	command := app.NewKubeletCommand()
	code := cli.Run(command)
	os.Exit(code)
}
