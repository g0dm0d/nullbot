package run

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Guild string `yaml:"guild,omitempty"`
	Token string `yaml:"token,omitempty"`
}

type Bot struct {
	session *discordgo.Session
	plugins []Plugin
}

func NewBot(token string) (*Bot, error) {
	s, err := discordgo.New("Bot " + token)
	return &Bot{session: s}, err
}

func Init() (Config, error) {
	var config Config
	configFile, err := os.ReadFile("config.yml")
	if err != nil {
		log.Panic(err)
	}
	err = yaml.Unmarshal([]byte(configFile), &config)
	return config, err
}

func (b *Bot) Run() {
	b.LoadHandlers()
	b.session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := b.session.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	b.LoadCommands()

	defer b.session.Close()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Println("Gracefully shutting down.")
}

type Plugin struct {
	Name     string
	Commands []Command
}

type Command struct {
	Command *discordgo.ApplicationCommand
	Func    func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func (b *Bot) RegisterPlugin(plugin *Plugin) {
	b.plugins = append(b.plugins, *plugin)
	log.Printf("Register plugin: %s", plugin.Name)
}

func (b *Bot) LoadCommands() {
	for _, p := range b.plugins {
		for _, c := range p.Commands {
			_, err := b.session.ApplicationCommandCreate(b.session.State.User.ID, "", c.Command)
			if err != nil {
				log.Panicf("Cannot create '%v' command: %v", c.Command.Name, err)
			} else {
				log.Printf("Register command: %s", p.Name)
			}
		}
	}
}

func (b *Bot) LoadHandlers() {
	for _, p := range b.plugins {
		for _, c := range p.Commands {
			b.session.AddHandler(c.Func)
		}
	}
}
