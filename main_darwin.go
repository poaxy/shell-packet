// +build darwin

package main

import (
	"fmt"
	"os/exec"
)

func configureDevice(name string) error {
	commands := [][]string{
		{"ifconfig", name, "inet", "10.0.0.1", "10.0.0.2", "up"},
		{"route", "add", "default", "10.0.0.2"},
	}

	for _, cmd := range commands {
		if err := exec.Command(cmd[0], cmd[1:]...).Run(); err != nil {
			return fmt.Errorf("failed to run command %v: %w", cmd, err)
		}
	}

	return nil
}
