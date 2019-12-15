package bot

import (
  "strings"

  "github.com/bwmarrin/discordgo"
)

func commandHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
  user := message.Author
	if user.ID == botID || user.Bot {
		return
	}

  channelID := message.ChannelID
  commandArgs := strings.Split(message.Content, " ")

  if !strings.HasPrefix(commandArgs[0], conf.Prefix) {
    return
  }

  command := commandArgs[0][len(conf.Prefix):]

  if commandArgs[0] == "!help" {
    handleHelp(session, channelID, commandArgs)

    return
  }

  if command == "hentai" {
    handleHentai(session, channelID, commandArgs)

    return
  }
}
