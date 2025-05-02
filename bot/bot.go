package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var Token string

func Run() {
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	discord.AddHandler(interactionHandler)

	err = discord.Open()
	if err != nil {
		log.Fatalf("Discord connection failed - %v", err)
	}
	defer discord.Close()

	fmt.Printf("Bot is running.  Press CTRL-C to exit.\n\nInvite the bot to your server with this link:\n  https://discord.com/api/oauth2/authorize?client_id=%v&scope=bot%%20applications.commands&permissions=105227086912\n\n", discord.State.User.ID)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	registerCommands(discord)
	<-sc
	unregisterCommands(discord)
}

func interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		log.Printf("Handling interaction '%v' in guild '%v'\n", i.ApplicationCommandData().Name, i.GuildID)
		switch i.ApplicationCommandData().Name {
		case "ping":
			handlePing(s, i)
		case "help":
			handleHelp(s, i)
		}
	}
}

func registerCommands(s *discordgo.Session) {
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
		log.Printf("Created command '%v'\n", v.Name)
	}
}

func unregisterCommands(s *discordgo.Session) {
	for _, v := range commands {
		s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		log.Printf("Deleted command '%v'\n", v.Name)
	}
}
