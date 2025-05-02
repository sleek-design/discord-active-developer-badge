package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "help",
		Description: "Links to the bot's documentation",
	},
	{
		Name:        "ping",
		Description: "Pings the bot to determine the latency",
	},
}

func handleHelp(s *discordgo.Session, i *discordgo.InteractionCreate) {
	url := "https://github.com/benjammin4dayz/discord-active-developer-badge?tab=readme-ov-file#getting-started"
	basicInteractionResponse(s, i,
		fmt.Sprintf("[Click here](%v) to learn how to use this bot to earn your Active Developer badge!", url))
}

func handlePing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	basicInteractionResponse(s, i,
		fmt.Sprintf("Pong! (%v ms)",
			s.HeartbeatLatency().Milliseconds()))
}

func basicInteractionResponse(s *discordgo.Session, i *discordgo.InteractionCreate, response string) {
	log.Println(response)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: response,
		},
	})
}
