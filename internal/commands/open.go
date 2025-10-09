package commands

import (
	"gitbroski/internal/services"
	"gitbroski/utils/git"
	"gitbroski/utils/logger"
)

func init() {
	Register("open", Open)
}

func Open(_ ...string) {
	logger.Text("Open command executed")
	url := git.GetRemoteURL()
	if url == "" {
		logger.Error("No URL found, aborting open command")
		return
	}
	logger.Text("Opening browser with URL: " + url)
	if err := services.OpenBrowser(url); err != nil {
		logger.Error("Failed to open browser: " + err.Error())
	}
}
