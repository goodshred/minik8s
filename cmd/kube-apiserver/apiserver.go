package kube_apiserver

import (
	"k8s.io/component-base/cli"
	"k8s.io/kubernetes/cmd/kube-apiserver/app"
	"os"
)

func main() {
	command := app.NewAPIServerCommand()
	code := cli.Run(command)
	os.Exit(code)
}
