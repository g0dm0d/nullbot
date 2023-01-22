package gcc

import (
	"log"
	"nullbot/plugins/compiler/tools"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func DiscordGcc(s *discordgo.Session, i *discordgo.InteractionCreate) {
	tools.GenerateRespond(runGcc, s, i)
}

func runGcc(code string) string {
	result, err := exec.Command("docker", "run", "-i", "--name", "script-gcc", "--rm", "gcc-compile",
		"/bin/bash", "-c", tools.GenerateCommand(tools.FixSymbol(code), "c")).CombinedOutput()
	if err != nil {
		log.Println(err)
		return strings.Join([]string{"err:", err.Error(), "\nmessage:", string(result)}, "")
	}
	return string(result)
}
