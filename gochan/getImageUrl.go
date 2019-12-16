package gochan

import (
  "strings"

  "github.com/Lol3rrr/dom"
)

func getImageUrl(tag dom.Tag) (string, bool) {
  rawUrl, ok := tag.Attributes["href"]
  if !ok {
    return "", false
  }

  url := strings.ReplaceAll(rawUrl, "//", "http://")

  return url, true
}
