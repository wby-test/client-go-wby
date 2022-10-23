package cluster

import (
	"fmt"
	"testing"
)

func TestGetConfigFile(t *testing.T) {
	out := GetConfigFile("kubetest")
	fmt.Println(out)
}
