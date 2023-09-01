# `terratest-kind`

This is a very simple library which allows you to build `kind` clusters for
the purposes of testing within `terratest`.

## Usage

```go
package test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/shell"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestBasicExample(t *testing.T) {
	cluster := cluster.New(t)

	cluster.Create()
	defer cluster.Delete()

	options := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/basic",
		EnvVars:      cluster.EnvVars(),
	})
	defer terraform.Destroy(t, options)

	terraform.InitAndApply(t, options)
}
```
