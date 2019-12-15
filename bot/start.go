package bot

import (
  "github.com/bwmarrin/discordgo"

  "hentai_bot/config"
  "hentai_bot/gochan"
)

var botID string
var conf *config.Config
var hentaiBoard gochan.Board

func Start(botToken string, pConf *config.Config) bool {
  conf = pConf

  discord, err := discordgo.New("Bot " + botToken)
  if err != nil {
    return false
  }

  user, err := discord.User("@me")
	if err != nil {
		return false
	}

  botID = user.ID
  discord.AddHandler(commandHandler)

  err = discord.Open()
  if err != nil {
    return false
  }

  go UpdateBoard()

  return true
}
