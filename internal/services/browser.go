// Package services provides core business logic services.
package services

import (
	"gitbroski/utils/logger"
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		logger.Error("unsupported platform")
	}
	return cmd.Start()
}
