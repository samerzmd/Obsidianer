package config

import (
	"testing"
)

type MockExecCommander struct{}

func (m MockExecCommander) RunCommand(command string, args ...string) ([]byte, error) {
	return []byte("mocked_username"), nil
}

func TestGetGitConfig(t *testing.T) {
	mock := MockExecCommander{}
	key := "user.name"
	prompt := "Please enter your Git username:"
	expected := "mocked_username"

	result, err := GetGitConfig(mock, key, prompt)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
