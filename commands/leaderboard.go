package commands

import (
	"fmt"
	"strings"

	"github.com/SharkBot-Game-Dev/CookieChan/consts"
	"github.com/SharkBot-Game-Dev/CookieChan/models"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var LeaderboardCommand = discord.SlashCommandCreate{
	Name:        "leaderboard",
	Description: "クッキー数の全体ランキングを上位10人まで表示します。",
}

func LeaderboardCommandExecute(interaction events.ApplicationCommandInteractionCreate, client bot.Client) bool {
	if err := client.Rest.CreateInteractionResponse(interaction.ID(), interaction.Token(), discord.InteractionResponse{
		Type: discord.InteractionResponseTypeDeferredCreateMessage,
	}); err != nil {
		return true
	}

	var users []models.CookieGameUser
	if err := consts.DB.Select("user_id", "cookie_count").
		Order("cookie_count DESC").Order("user_id ASC").Limit(10).Find(&users).Error; err != nil {
		client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), discord.MessageCreate{
			Content: "内部DBエラーが発生しました。",
		})
		return true
	}

	var description strings.Builder
	for i, user := range users {
		fmt.Fprintf(&description, "%d位：<@%s> — 🍪 %dクッキー\n", i+1, user.UserId, user.CookieCount)
	}
	if len(users) == 0 {
		description.WriteString("まだランキングに参加しているユーザーはいません。")
	}

	client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), discord.MessageCreate{
		Embeds: []discord.Embed{{
			Title:       "🏆クッキー数ランキング",
			Description: description.String(),
			Color:       consts.CookieColor,
			Footer: &discord.EmbedFooter{
				Text: "上位10位を表示しています。",
			},
		}},
	})
	return true
}
