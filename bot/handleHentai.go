package bot

import (
  "github.com/bwmarrin/discordgo"
)

func handleHentai(session *discordgo.Session, channelID string, parts []string) {
  post, worked := hentaiBoard.SelectRandomPost()
  if !worked {
    session.ChannelMessageSend(channelID, "Could not select a random post")
  }

  message := post.Image

  session.ChannelMessageSend(channelID, message)
}
