package python

import (
	"log"
	"nullbot/tools"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func DiscordPython(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
						Reader:      strings.NewReader(runPython(tools.GetFile(test.URL))),
					},
				},
			},
		})
	}
}

func runPython(code string) string {
	result, err := exec.Command("sh", "./tools/python/runner.sh", code).CombinedOutput()
	if err != nil {
		log.Println(err)
		return strings.Join([]string{"err:", err.Error(), "\nmessage:", string(result)}, "")
	}
	return string(result)
}
