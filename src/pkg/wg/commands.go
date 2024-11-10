package wg

import (
	"bytes"
	"os/exec"
)

func RunCommandWithStoredPassword(password string, command []string) (string, error) {
	cmd := exec.Command("sudo", append([]string{"-S"}, command...)...)
	cmd.Stdin = bytes.NewBufferString(password + "\n")
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	err := cmd.Run()
	return output.String(), err
}
