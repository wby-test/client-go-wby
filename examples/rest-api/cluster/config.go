package cluster

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GetConfigFile(env string) string {
	cmd := exec.Command("ls", "-al")
	outInfo := bytes.Buffer{}
	cmd.Stdout = &outInfo
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}

	return outInfo.String()
}
