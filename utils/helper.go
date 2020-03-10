package utils

import (
	"os"
	"os/exec"
)

func RunCommandWithOutput(c *exec.Cmd) (string, error) {
	c.Stderr = os.Stderr

	data, err := c.Output()
	if err != nil {
		return "", err
	}

	return string(data), nil
}
