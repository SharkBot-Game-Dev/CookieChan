package commands

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
)

type CommandExecute func(discord.Interaction, bot.Client) bool

var Commands = []discord.SlashCommandCreate{}
var CommandExecutes = make(map[string]CommandExecute)

func InitCommand() {
	Commands = append(Commands, HelpCommand)
	CommandExecutes[HelpCommand.Name] = HelpCommandExecute

	Commands = append(Commands, JankenCommand)
	CommandExecutes[JankenCommand.Name] = JankenCommandExecute
}
