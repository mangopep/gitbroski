package git

import (
	"gitbroski/utils/logger"
	"os/exec"
	"strings"
)

func GetRoot() string {

	rootCmd := exec.Command("git", "rev-parse", "--show-toplevel")
	rootOut, err := rootCmd.Output()

	if err != nil {
		logger.Error("Not inside a git repository")
		return ""
	}
	rootDir := strings.TrimSpace(string(rootOut))

	return rootDir
}
