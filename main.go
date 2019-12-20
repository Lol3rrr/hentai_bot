package main

import (
  "fmt"
  "time"

  "hentai_bot/bot"
  "hentai_bot/config"
  "hentai_bot/general"
)

func main() {
  fmt.Printf("Starting... \n")

  configFile := "config.json"

  fmt.Printf("Loading Config... \n")
  botConfig, err := config.Load(configFile)
  if general.CheckError(err, "Main") {
    fmt.Printf("Error loading Config: '%v' \n", err)
    return
  }

  fmt.Printf("Starting Bot... \n")
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
