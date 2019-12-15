package main

import (
  "fmt"
  "time"

  "hentai_bot/bot"
  "hentai_bot/config"
  "hentai_bot/general"
)

func main() {
  configFile := "./config.json"

  botConfig, err := config.Load(configFile)
  if general.CheckError(err, "Main") {
    return
  }

  started := bot.Start(botConfig.BotToken, botConfig)
  if !started {
    fmt.Printf("Error starting bot \n")
    return
  }
  fmt.Printf("Started Bot \n")

  for {
    time.Sleep(10000)
  }
}
