package gochan

import (
  "github.com/Lol3rrr/dom"
)

func getPageThreads(boardPage dom.TagList) []string {
  result := make([]string, 0)

  threads := boardPage.SearchAttribute("class", "thread")
  for _, thread := range threads {
    rawID, found := thread.Attributes["id"]
    if !found || len(rawID) < 2 {
      continue
    }
    id := rawID[1:]

    result = append(result, id)
  }

  return result
}
