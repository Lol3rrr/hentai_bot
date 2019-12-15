package bot

import (
  "time"

  "hentai_bot/gochan"
)

func UpdateBoard() {
  hentaiBoard = gochan.LoadBoard("h")

  for {
    time.Sleep(10 * time.Minute)

    hentaiBoard.Refresh()
  }
}
