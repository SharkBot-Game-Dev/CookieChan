package commands

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

type CommandExecute func(events.ApplicationCommandInteractionCreate, bot.Client) bool

var Commands = []discord.SlashCommandCreate{}
var CommandExecutes = make(map[string]CommandExecute)

func InitCommand() {
	Commands = nil
	CommandExecutes = make(map[string]CommandExecute)
	Commands = append(Commands, HelpCommand)
	CommandExecutes[HelpCommand.Name] = HelpCommandExecute

	// Commands = append(Commands, JankenCommand)
	// CommandExecutes[JankenCommand.Name] = JankenCommandExecute

	Commands = append(Commands, CookieCommand)
	CommandExecutes[CookieCommand.Name] = CookieCommandExecute

	Commands = append(Commands, LeaderboardCommand)
	CommandExecutes[LeaderboardCommand.Name] = LeaderboardCommandExecute

	Commands = append(Commands, BuyCommand)
	CommandExecutes[BuyCommand.Name] = BuyCommandExecute

	Commands = append(Commands, ShopCommand)
	CommandExecutes[ShopCommand.Name] = ShopCommandExecute

	Commands = append(Commands, ClickCommand)
	CommandExecutes[ClickCommand.Name] = ClickCommandExecute

	Commands = append(Commands, VoteCommand)
	CommandExecutes[VoteCommand.Name] = VoteCommandExecute
}
