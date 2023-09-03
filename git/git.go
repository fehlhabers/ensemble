package git

import (
	"fmt"
	"os/exec"
)

var (
	mainBranch = "main"
)

func Pull() error {
	return exec.Command("git", "pull").Run()
}

func NewBranch(branch string, message string) error {
	if err := exec.Command("git", "checkout", "-b", branch).Run(); err != nil {
		return err
	}

	if err := exec.Command("git", "commit", "-m", message, "--allow-empty").Run(); err != nil {
		return err
	}

	return exec.Command("git", "push", "-u", "origin", fmt.Sprintf("%s/%s", mainBranch, branch)).Run()
}
