package bot

import (
  "hentai_bot/config"
)

func getHelpMessage(tmpConf *config.Config) string {
  message := "```"

  message += "Prefix: '"
  message += tmpConf.Prefix
  message += "'\n"
  message += "\n"

  message += "Commands: \n"
  message += "  - hentai: Sends a random post from the /h/ Board"

  message += "```"

  return message
}
