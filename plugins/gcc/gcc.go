package gcc

import (
	"log"
	"nullbot/plugins"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func DiscordGcc(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}
	if option, ok := optionMap["code"]; ok {
		test := i.Interaction.ApplicationCommandData().Resolved.Attachments[option.Value.(string)]
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You're code return:",
				Files: []*discordgo.File{
					{
						ContentType: "text/plain",
						Name:        "output.txt",
						Reader:      strings.NewReader(runGcc(plugins.GetFile(test.URL))),
					},
				},
			},
		})
	}
}

func runGcc(code string) string {
	result, err := exec.Command("docker", "run", "-i", "--name", "script-gcc", "--rm", "gcc-compile",
		"/bin/bash", "-c", plugins.GenerateCommand(plugins.FixSymbol(code), "c")).CombinedOutput()
	if err != nil {
		log.Println(err)
		return strings.Join([]string{"err:", err.Error(), "\nmessage:", string(result)}, "")
	}
	return string(result)
}
