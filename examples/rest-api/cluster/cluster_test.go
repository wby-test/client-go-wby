package cluster

import (
	"fmt"
	"testing"
)

func TestGetConfigFile(t *testing.T) {
	out := GetConfigFile("tes")
	fmt.Println(out)
}
