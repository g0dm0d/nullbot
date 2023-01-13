package main

import (
	"flag"
	"log"
	"nullbot/tools/gcc"
	"nullbot/tools/python"
	"nullbot/tools/rust"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

// Bot parameters
var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", "", "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", true, "")
)

var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

var (
	integerOptionMinValue          = 1.0
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageServer

	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "python",
			Description: "run python script",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionAttachment,
					Name:        "code",
					Description: "drop code file",
					Required:    true,
				},
			},
		},
		{
			Name:        "rust",
			Description: "compile rust code",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionAttachment,
					Name:        "code",
					Description: "drop code file",
					Required:    true,
				},
			},
		},
		{
			Name:        "gcc",
			Description: "compile C/C++ code",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionAttachment,
					Name:        "code",
					Description: "drop code file",
					Required:    true,
				},
			},
		},
		{
			Name:        "followups",
			Description: "Followup messages",
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"python": python.DiscordPython,
		"rust":   rust.DiscordRust,
		"gcc":    gcc.DiscordGcc,
	}
)

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if *RemoveCommands {
		log.Println("Removing commands...")
		// // We need to fetch the commands, since deleting requires the command ID.
		// // We are doing this from the returned commands on line 375, because using
		// // this will delete all the commands, which might not be desirable, so we
		// // are deleting only the commands that we added.
		// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)
		// if err != nil {
		// 	log.Fatalf("Could not fetch registered commands: %v", err)
		// }

		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, *GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")
}

// func main() {
// 	test := "print(\"hello world!\")"
// 	a, err := exec.Command("sh", "/home/godmod/Documents/nullbot/tools/python/runner.sh", test).Output()
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println("err")
// 	}
// 	fmt.Println(string(a))
// }
