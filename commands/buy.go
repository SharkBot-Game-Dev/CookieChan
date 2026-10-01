package commands

import (
	"errors"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"gorm.io/gorm"

	"github.com/SharkBot-Game-Dev/CookieChan/consts"
	"github.com/SharkBot-Game-Dev/CookieChan/models"

	"strconv"

	"time"
)

var BuyCommand = discord.SlashCommandCreate{
	Name:        "buy",
	Description: "アイテムをクッキーを使って購入します。",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionString{
			Name:        "item",
			Description: "購入するアイテム名",
			Required:    true,
		},
		discord.ApplicationCommandOptionInt{
			Name:        "count",
			Description: "購入するアイテムの個数",
			Required:    false,
		},
	},
}

func BuyCommandExecute(interaction events.ApplicationCommandInteractionCreate, client bot.Client) bool {
	client.Rest.CreateInteractionResponse(interaction.ID(), interaction.Token(), discord.InteractionResponse{
		Type: discord.InteractionResponseTypeDeferredCreateMessage,
	})

	var cookieUser models.CookieGameUser

	itemName := interaction.SlashCommandInteractionData().String("item")
	count, countOK := interaction.SlashCommandInteractionData().OptInt("count")
	if !countOK {
		count = 1
	}

	if count < 0 {
		client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), discord.MessageCreate{Content: "1以上を指定してください。"})
		return true
	}

	if count > 50 {
		client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), discord.MessageCreate{Content: "一度には50個までしか買えません。"})
		return true
	}

	item := consts.ItemNameToStruct(itemName)
	if item == nil {
		client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), discord.MessageCreate{Content: "そのアイテムはありません。\n/shopコマンドで再度確認してください。"})
		return true
	}

	cookieUserResult := consts.DB.First(cookieUser, "user_id = ?", interaction.User().ID.String())
	if errors.Is(cookieUserResult.Error, gorm.ErrRecordNotFound) {
		if cookieUserResult.Error != nil {
			client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), discord.MessageCreate{Content: "まだクッキーがありません。\nまずは、/clickを実行してみましょう！"})
			return true
		}
	} else {
		if cookieUserResult.Error != nil {
			client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), discord.MessageCreate{Content: "内部DBエラーが発生しました。"})
			return true
		}
	}

	if cookieUser.CookieCount < (item.Price * count) {
		client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), discord.MessageCreate{Content: "お金が足りません。\nあなたは" + strconv.Itoa(cookieUser.CookieCount) + "クッキーまでのアイテムを購入できます。"})
		return true
	}

	updated := false
	nowItems := cookieUser.Items

	for itemI, itemNow := range nowItems {
		if itemNow.ItemId == item.ID {
			itemCopyNow := itemNow
			itemCopyNow.Count += count
			nowItems[itemI] = itemCopyNow
			updated = true
			break
		}
	}

	if !updated {
		nowItems = append(nowItems, models.CookieGameItem{
			UserId:    interaction.User().ID.String(),
			ItemId:    item.ID,
			Count:     count,
			CreatedAt: time.Now(),
		})
	}

	consts.DB.Model(&cookieUserResult).Update("items", nowItems)

	helpEmbed := discord.Embed{
		Title: "アイテムを購入しました。",
		Color: consts.CookieColor,
		Fields: []discord.EmbedField{
			{
				Name:   "🍮アイテム名",
				Value:  item.Name,
				Inline: &consts.False,
			},
			{
				Name:   "🍪値段",
				Value:  strconv.Itoa(item.Price*count) + "クッキー",
				Inline: &consts.False,
			},
		},
	}

	message := discord.MessageCreate{Embeds: []discord.Embed{helpEmbed}}

	client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), message)
	// log.Print(err)

	return true
}
