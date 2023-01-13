package gcc

import (
	"log"
	"nullbot/tools"
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
		log.Println(tools.GetFile(test.URL))
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You're code return:",
				Files: []*discordgo.File{
					{
						ContentType: "text/plain",
						Name:        "output.txt",
						Reader:      strings.NewReader(runGcc(tools.GetFile(test.URL))),
					},
				},
			},
		})
	}
}

func runGcc(code string) string {
	result, err := exec.Command("docker", "run", "-i", "--rm", "--name", "script-gcc", "gcc-compile", "sh", "compiler.sh", "|", "echo", code).Output()
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	return string(result)
}
