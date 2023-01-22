package rust

import (
	"log"
	"nullbot/plugins/compiler/tools"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func DiscordRust(s *discordgo.Session, i *discordgo.InteractionCreate) {
	tools.GenerateRespond(runRust, s, i)
}

func runRust(code string) string {
	result, err := exec.Command("docker", "run", "-i", "--name", "script-rust", "--rm", "rust-compile",
		"/bin/bash", "-c", tools.GenerateCommand(tools.FixSymbol(code), "rs")).CombinedOutput()
	if err != nil {
		log.Println(err)
		return strings.Join([]string{"err:", err.Error(), "\nmessage:", string(result)}, "")
	}
	return string(result)
}
