// Package git provides utilities for interacting with git repositories.
package git

import (
	"gitbroski/utils/logger"
	"os/exec"
	"regexp"
	"strings"
)

func GetRemoteURL() string {
	remoteCmd := exec.Command("git", "remote", "-v")
	remoteOut, err := remoteCmd.Output()

	re := regexp.MustCompile(`https?://[^\s]+`)
	URL := re.Find(remoteOut)

	if err != nil {
		logger.Error("Failed to get remote URL")
		return ""
	}
	remoteURL := strings.TrimSpace(string(URL))
	logger.Text("Remote URL: " + remoteURL)
	return remoteURL
}
