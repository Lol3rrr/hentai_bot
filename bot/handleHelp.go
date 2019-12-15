package bot

import (
  "github.com/bwmarrin/discordgo"
)

func handleHelp(session *discordgo.Session, channelID string, parts []string) {
  message := getHelpMessage(conf)

  session.ChannelMessageSend(channelID, message)
}
