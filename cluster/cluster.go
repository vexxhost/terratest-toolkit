package cluster

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/shell"
)

type Cluster struct {
	Name string
	*testing.T
}

func New(t *testing.T) *Cluster {
	return &Cluster{
		Name: strings.ToLower(random.UniqueId()),
		T:    t,
	}
}

func (k *Cluster) Create() {
	shell.RunCommand(k.T, shell.Command{
		Command: "kind",
		Args:    []string{"create", "cluster", "--name", k.Name},
	})
}

func (k *Cluster) Delete() {
	shell.RunCommand(k.T, shell.Command{
		Command: "kind",
		Args:    []string{"delete", "cluster", "--name", k.Name},
	})
}

func (k *Cluster) EnvVars() map[string]string {
	return map[string]string{
		"KUBE_CONFIG_PATH": fmt.Sprintf("%s/.kube/config", os.Getenv("HOME")),
		"KUBE_CTX":         fmt.Sprintf("kind-%s", k.Name),
	}
}
