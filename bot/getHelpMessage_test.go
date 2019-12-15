package bot

import (
  "testing"

  "hentai_bot/config"
)

func TestGetHelpMessage(t *testing.T) {
  tables := []struct{
    InputConfig *config.Config
    ResultString string
  }{
    {
      InputConfig: &config.Config{},
      ResultString: "```Prefix: ''\n\nCommands: \n  - hentai: Sends a random post from the /h/ Board```",
    },
    {
      InputConfig: &config.Config{
        Prefix: "!",
      },
      ResultString: "```Prefix: '!'\n\nCommands: \n  - hentai: Sends a random post from the /h/ Board```",
    },
  }

  for _, table := range tables {
    inConfig := table.InputConfig
    resString := table.ResultString

    outputString := getHelpMessage(inConfig)

    if outputString != resString {
      t.Errorf("Did not return the correct string, returned '%s', expected '%s' \n", outputString, resString)
      continue
    }
  }
}
