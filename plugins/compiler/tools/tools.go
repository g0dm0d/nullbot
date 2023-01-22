package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func GetFile(file string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", file, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(bodyText)
}

func GenerateRespond(runner func(code string) string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}
	if option, ok := optionMap["code"]; ok {
		file := i.Interaction.ApplicationCommandData().Resolved.Attachments[option.Value.(string)]
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You're code return:",
				Files: []*discordgo.File{
					{
						ContentType: "text/plain",
						Name:        "output.txt",
						Reader:      strings.NewReader(runner(GetFile(file.URL))),
					},
				},
			},
		})
	}
}

func FixSymbol(code string) string {
	return strings.Replace(code, "\"", "\\\"", -1)
}

func GenerateCommand(code, filetype string) string {
	return fmt.Sprintf("echo \"%s\" > file.%s && sh compiler.sh", code, filetype)
}
