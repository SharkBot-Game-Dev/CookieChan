package commands

import (
	"math/rand"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"

	"github.com/SharkBot-Game-Dev/CookieChan/cache"
	"github.com/SharkBot-Game-Dev/CookieChan/consts"
)

var JankenCommand = discord.SlashCommandCreate{
	Name:        "janken",
	Description: "私とじゃんけんするよ！",
}

func JankenCommandExecute(interaction events.ApplicationCommandInteractionCreate, client bot.Client) bool {
	if err := client.Rest.CreateInteractionResponse(interaction.ID(), interaction.Token(), discord.InteractionResponse{
		Type: discord.InteractionResponseTypeDeferredCreateMessage,
	}); err != nil {
		return true
	}

	answers := []string{"ぐー", "ちょき", "ぱー"}
	if !cache.StartJanken(interaction.User().ID, interaction.Channel().ID(), answers[rand.Intn(len(answers))]) {
		client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), discord.MessageCreate{Content: "ゲームを始めたばっかりだよ！\nぐー、ちょき、ぱーのどれかを入力してね！"})
		return true
	}

	helpEmbed := discord.Embed{
		Title:       "じゃんけん！",
		Description: "ぐー、ちょき、ぱーのどれかを入力してね！",
		Color:       consts.CookieColor,
	}

	message := discord.MessageCreate{Embeds: []discord.Embed{helpEmbed}}

	client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), message)
	// log.Print(err)

	return true
}
