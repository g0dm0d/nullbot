package python

import (
	"log"
	"nullbot/plugins/compiler/tools"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func DiscordPython(s *discordgo.Session, i *discordgo.InteractionCreate) {
	tools.GenerateRespond(runPython, s, i)
}

func runPython(code string) string {
	result, err := exec.Command("sh", "./plugins/compiler/python/runner.sh", code).CombinedOutput()
	if err != nil {
		log.Println(err)
		return strings.Join([]string{"err:", err.Error(), "\nmessage:", string(result)}, "")
	}
	return string(result)
}
