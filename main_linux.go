// +build linux

package main

import (
	"fmt"
	"os/exec"
)

func configureDevice(name string) error {
	commands := [][]string{
		{"ip", "addr", "add", "10.0.0.1/24", "dev", name},
		{"ip", "link", "set", "up", "dev", name},
		{"ip", "route", "add", "default", "via", "10.0.0.1", "dev", name},
	}

	for _, cmd := range commands {
		if err := exec.Command(cmd[0], cmd[1:]...).Run(); err != nil {
			return fmt.Errorf("failed to run command %v: %w", cmd, err)
		}
	}

	return nil
}
