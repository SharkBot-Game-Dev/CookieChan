package commands

import (
	"strconv"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"

	"github.com/SharkBot-Game-Dev/CookieChan/consts"
)

var ClickCommand = discord.SlashCommandCreate{
	Name:        "click",
	Description: "クッキーをクリックします。",
}

func ClickCommandExecute(
	interaction events.ApplicationCommandInteractionCreate,
	client bot.Client,
) bool {
	if err := client.Rest.CreateInteractionResponse(
		interaction.ID(),
		interaction.Token(),
		discord.InteractionResponse{
			Type: discord.InteractionResponseTypeDeferredCreateMessage,
		},
	); err != nil {
		return true
	}

	userID := interaction.User().ID.String()

	cookieUser, clickCount, err := clickCookies(consts.DB, userID)
	if err != nil {
		client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), discord.MessageCreate{Content: "内部DBエラーが発生しました。"})
		return true
	}
	cookieCount := cookieUser.CookieCount

	clickEmbed := discord.Embed{
		Title: "🖱️クッキーをクリックしました！",
		Description: strconv.Itoa(clickCount) +
			"クッキーを入手しました。\n現在は" +
			strconv.Itoa(cookieCount) +
			"クッキー持っています。",
		Color: consts.CookieColor,
	}

	client.Rest.CreateFollowupMessage(
		client.ApplicationID,
		interaction.Token(),
		discord.MessageCreate{
			Embeds: []discord.Embed{clickEmbed},
		},
	)

	for _, achievement := range consts.CookieAchievements {
		if achievement.CookieCount > cookieCount {
			continue
		}
		unlocked, err := unlockAchievement(consts.DB, userID, achievement.ID)
		if err != nil || !unlocked {
			continue
		}
		client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), discord.MessageCreate{Content: "やったー！新しい実績を達成したよ！\n「" + achievement.Name + "」\n/cookieで確認してみてね！"})
	}
	return true
}
