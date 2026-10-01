package commands

import (
	"github.com/SharkBot-Game-Dev/CookieChan/consts"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/omit"
)

var permission = omit.NewPtr(discord.PermissionManageChannels)

var VoteCommand = discord.SlashCommandCreate{
	Name:                     "vote",
	Description:              "アンケートを作成します。",
	DefaultMemberPermissions: permission,
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionString{
			Name:        "title",
			Description: "アンケートのタイトル",
			Required:    true,
		},
		discord.ApplicationCommandOptionString{
			Name:        "option_1",
			Description: "オプション１",
			Required:    true,
		},
		discord.ApplicationCommandOptionString{
			Name:        "option_2",
			Description: "オプション２",
			Required:    false,
		},
		discord.ApplicationCommandOptionString{
			Name:        "option_3",
			Description: "オプション３",
			Required:    false,
		},
		discord.ApplicationCommandOptionString{
			Name:        "option_4",
			Description: "オプション４",
			Required:    false,
		},
		discord.ApplicationCommandOptionString{
			Name:        "option_5",
			Description: "オプション５",
			Required:    false,
		},
	},
}

var number_emoji_map = map[int]string{
	1: "1️⃣",
	2: "2️⃣",
	3: "3️⃣",
	4: "4️⃣",
	5: "5️⃣",
}

func VoteCommandExecute(interaction events.ApplicationCommandInteractionCreate, client bot.Client) bool {
	if err := client.Rest.CreateInteractionResponse(interaction.ID(), interaction.Token(), discord.InteractionResponse{
		Type: discord.InteractionResponseTypeDeferredCreateMessage,
	}); err != nil {
		return true
	}

	title := interaction.SlashCommandInteractionData().String("title")
	options := []string{}

	option_1 := interaction.SlashCommandInteractionData().String("option_1")
	options = append(options, option_1)

	option_2, option_2_ok := interaction.SlashCommandInteractionData().OptString("option_2")
	if option_2_ok {
		options = append(options, option_2)
	}

	option_3, option_3_ok := interaction.SlashCommandInteractionData().OptString("option_3")
	if option_3_ok {
		options = append(options, option_3)
	}

	option_4, option_4_ok := interaction.SlashCommandInteractionData().OptString("option_4")
	if option_4_ok {
		options = append(options, option_4)
	}

	option_5, option_5_ok := interaction.SlashCommandInteractionData().OptString("option_5")
	if option_5_ok {
		options = append(options, option_5)
	}

	description := ""
	for count, option := range options {
		if len(options) == 2 {
			emoji := ""
			if count == 0 {
				emoji = "🅰️"
			} else {
				emoji = "🇧"
			}

			description += emoji + " " + option + "\n"
		} else if len(options) == 1 {
			description += "👍" + " " + option + "\n"
		} else {
			description += number_emoji_map[count+1] + " " + option + "\n"
		}
	}

	helpEmbed := discord.Embed{
		Title:       title,
		Color:       consts.CookieColor,
		Description: description,
	}

	message := discord.MessageCreate{Embeds: []discord.Embed{helpEmbed}}

	followUpMessage, followUpErr := client.Rest.CreateFollowupMessage(client.ApplicationID, interaction.Token(), message)
	// log.Print(err)
	if followUpErr != nil {
		return true
	}

	if len(options) == 2 {
		client.Rest.AddReaction(interaction.Channel().ID(), followUpMessage.ID, "🅰️")
		client.Rest.AddReaction(interaction.Channel().ID(), followUpMessage.ID, "🇧")
	} else if len(options) == 1 {
		client.Rest.AddReaction(interaction.Channel().ID(), followUpMessage.ID, "👍")
	} else {
		for count, _ := range options {
			client.Rest.AddReaction(interaction.Channel().ID(), followUpMessage.ID, number_emoji_map[count+1])
		}
	}

	return true
}
