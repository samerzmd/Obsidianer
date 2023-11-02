package config

import (
	"fmt"
	"os/exec"
	"strings"
)

type ExecCommander interface {
	RunCommand(command string, args ...string) ([]byte, error)
}

type RealExecCommander struct{}

func (r RealExecCommander) RunCommand(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).Output()
}

func GetGitConfig(execCommander ExecCommander, key string, prompt string) (string, error) {
	output, err := execCommander.RunCommand("git", "config", "--global", key)
	if err != nil || len(output) == 0 {
		fmt.Println(prompt)
		var userInput string
		fmt.Scanln(&userInput)
		_, err := execCommander.RunCommand("git", "config", "--global", key, userInput)
		if err != nil {
			return "", err
		}
		return userInput, nil
	}
	return strings.TrimSpace(string(output)), nil
}
