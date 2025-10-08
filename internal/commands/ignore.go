// Package commands provides CLI command handlers.
package commands

import (
	"gitbroski/internal/services"
	"gitbroski/utils/logger"
)

func init() {
	Register("ignore", Ignore)
}

func Ignore(args ...string) {
	if len(args) == 0 {
		logger.Error("No arguments provided to ignore command....Generating empty")
		services.Ignore("")
		return
	}
	services.Ignore(args[0])
}
