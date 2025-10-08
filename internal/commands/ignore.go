package commands

import (
	"gitbroski/internal/services"
	"gitbroski/utils/logger"
	"strings"
)

func init() {
	Register("ignore", Ignore)
}

func Ignore(args ...string) {
	if len(args) == 0 {
		logger.Warning("⚠️  No language specified. Usage: gitbroski ignore <language>")
		logger.Text("📋 Supported languages:")
		logger.Text("   • python")
		logger.Text("   • node, nodejs, js, javascript")
		logger.Text("\n💡 Example: gitbroski ignore node")
		return
	}
	
	lang := strings.ToLower(strings.TrimSpace(args[0]))
	if lang == "" {
		logger.Error("❌ Invalid language specified")
		return
	}
	
	services.Ignore(lang)
}
