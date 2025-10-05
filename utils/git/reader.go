package git

import (
	"gitbroski/utils/logger"
	"os/exec"
	"strings"
)

func GetRemoteURL() string {
	// Getting the root of the git repository
	rootCmd := exec.Command("git", "rev-parse", "--show-toplevel")
	rootOut, err := rootCmd.Output()
	if err != nil {
		logger.Error("Not inside a git repository")
		return ""
	}
	repoRoot := strings.TrimSpace(string(rootOut))

	// Getting remote URL from the repo root
	remoteCmd := exec.Command("git", "-C", repoRoot, "config", "--get", "remote.origin.url")
	remoteOut, err := remoteCmd.Output()
	if err != nil {
		logger.Error("Failed to get remote URL")
		return ""
	}
	remoteURL := strings.TrimSpace(string(remoteOut))
	logger.Text("Remote URL: " + remoteURL)
	return remoteURL
}
