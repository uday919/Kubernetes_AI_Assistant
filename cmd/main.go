package main

import (
	"github.com/uday919/Kubernetes_AI_Assistant/cmd/cli"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	cli.InitAndExecute()
}
