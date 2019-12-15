package bot

import (
  "github.com/bwmarrin/discordgo"
)

func handleHelp(session *discordgo.Session, channelID string, parts []string) {
  message := "```Prefix: '!'\n\nCommands: \n  - hentai: Sends a random post from the /h/ board```"

  session.ChannelMessageSend(channelID, message)
}
