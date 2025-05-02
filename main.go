package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	bot "github.com/benjammin4dayz/discord-active-developer-badge/bot"
	"github.com/joho/godotenv"
)

var BotToken = flag.String("token", "", "Bot access token")

func main() {
	godotenv.Load()

	if os.Getenv("DEBUG") != "true" {
		log.SetOutput(io.Discard)
	}

	bot.Token = getToken()
	bot.Run()
}

func getToken() string {
	flag.Parse()
	if *BotToken != "" {
		return *BotToken
	}

	token := os.Getenv("BOT_TOKEN")
	if token != "" {
		return token
	}

	fmt.Print("Enter bot token: ")
	fmt.Scanln(&token)
	return token
}
