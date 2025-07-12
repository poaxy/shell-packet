// +build windows

package main

import (
	"fmt"
	"os/exec"
)

func configureDevice(name string) error {
	commands := [][]string{
		{"netsh", "interface", "ip", "set", "address", fmt.Sprintf("name=\"%s\"", name), "static", "10.0.0.1", "255.255.255.0"},
		{"netsh", "interface", "ip", "set", "dns", fmt.Sprintf("name=\"%s\"", name), "static", "8.8.8.8"},
		{"netsh", "interface", "ip", "add", "route", "0.0.0.0/0", fmt.Sprintf("interface=\"%s\"", name), "10.0.0.1"},
	}

	for _, cmd := range commands {
		if err := exec.Command(cmd[0], cmd[1:]...).Run(); err != nil {
			return fmt.Errorf("failed to run command %v: %w", cmd, err)
		}
	}

	return nil
}
