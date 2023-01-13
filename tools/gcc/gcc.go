package gcc

import (
	"fmt"
	"io/ioutil"
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
	ioutil.WriteFile("/home/godmod/Documents/nullbot/tmp/file.c", []byte(code), 0644)
	result, err := exec.Command("docker", "run", "-i", "--name", "script-gcc", "--rm", "gcc-compile",
		"/bin/bash", "-c", fmt.Sprintf("echo \"%s\" > file.c && sh compiler.sh",
			strings.Replace(code, "\"", "\\\"", -1))).CombinedOutput()
	log.Println(string(result))
	if err != nil {
		log.Println(err)
		return strings.Join([]string{"err:", err.Error(), "\nmessage:", string(result)}, "")
	}
	return string(result)
}
