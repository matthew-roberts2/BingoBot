package bot

import (
	"bingoBotGo/internal/command"
	"bingoBotGo/internal/util"
	"github.com/bwmarrin/discordgo"
	"log"
)

const GUILD_NAME_EXPIRIY_IN_SECONDS = 60 * 60 * 24

type Bot struct {
	HumanName          string
	registeredCommands []command.Command
	session            *discordgo.Session
	guildNameCache     util.Cache[string]
}

func MakeBot(client *discordgo.Session, name string) Bot {
	return Bot{
		registeredCommands: []command.Command{},
		session:            client,
		HumanName:          name,
		guildNameCache:     util.MakeCache[string](),
	}
}

func (bot *Bot) RegisterCommand(command command.Command) {
	log.Println("Registering command ", command)
	bot.registeredCommands = append(bot.registeredCommands, command)
}

func (bot *Bot) GetCommandCount() int {
	return len(bot.registeredCommands)
}

func (bot Bot) IsSelf(userId string) bool {
	return userId == bot.Session().State.User.ID
}

func (bot Bot) GetInternalName() string {
	return bot.HumanName
}

func (bot Bot) Session() *discordgo.Session {
	return bot.session
}

func (bot Bot) GetGuildName(guildId string) string {
	name, exists := bot.guildNameCache.Get(guildId)
	if !exists {
		name = bot.lookupBotNameInGuild(guildId)
		bot.guildNameCache.Put(guildId, name, GUILD_NAME_EXPIRIY_IN_SECONDS)
	}
	return name
}

func (bot Bot) lookupBotNameInGuild(guildId string) string {
	member, err := bot.Session().GuildMember(guildId, bot.Session().State.User.ID)
	if err != nil {
		log.Printf("Failed to look up guild name for bot in guild id %s\n", guildId)
	}
	return member.Nick
}

func (bot Bot) HandleMessage(inboundMessage *discordgo.MessageCreate) {
	message := inboundMessage.Message
	var commandResult command.Result
	for _, botCommand := range bot.registeredCommands {
		commandResult = botCommand.Process(bot, message)
		if commandResult == command.SUCCESS {
			break
		}
	}

	if commandResult != command.SUCCESS && !bot.IsSelf(message.Author.ID) {
		log.Println("Bot handler failed to find any matching commands")
	}
}
