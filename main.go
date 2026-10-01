package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/snowflake/v2"
	"github.com/joho/godotenv"

	"github.com/SharkBot-Game-Dev/CookieChan/commands"
	"github.com/SharkBot-Game-Dev/CookieChan/consts"
	"github.com/SharkBot-Game-Dev/CookieChan/events"
)

var DISCORD_TOKEN = ""
var DISCORD_CLIENT_ID snowflake.ID
var SYNC_SLASH = false
var DSN = ""

func Env_load() {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error loading .env file")
	}

	DISCORD_TOKEN = os.Getenv("DISCORD_TOKEN")
	DISCORD_CLIENT_ID, err = snowflake.Parse(os.Getenv("DISCORD_CLIENT_ID"))
	if err != nil {
		log.Fatal("Error parsing DISCORD_CLIENT_ID")
	}
	SYNC_SLASH = os.Getenv("SYNC_SLASH") == "1"

	DSN = os.Getenv("DSN")
	if DSN == "" {
		log.Fatal("Error not found DSN")
	}

	consts.ConnectDB(DSN)
}

func main() {
	Env_load()

	client, err := disgo.New(DISCORD_TOKEN,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuilds,
				gateway.IntentGuildMessages,
				gateway.IntentMessageContent,
			),
		),

		bot.WithEventListenerFunc(events.MessageCreate),
		bot.WithEventListenerFunc(events.ApplicationCommandInteractionCreate),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close(context.Background())
	commands.InitCommand()

	if SYNC_SLASH {
		applicationCommands := make([]discord.ApplicationCommandCreate, len(commands.Commands))
		for i, command := range commands.Commands {
			applicationCommands[i] = command
		}
		if _, err := client.Rest.SetGlobalCommands(DISCORD_CLIENT_ID, applicationCommands); err != nil {
			log.Fatal(err)
		}

		log.Print("スラッシュコマンドを同期しました。")
	}

	if err = client.OpenGateway(context.TODO()); err != nil {
		panic(err)
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
}
