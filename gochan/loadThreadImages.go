package gochan

import (
  "strings"

  "hentai_bot/dom"
)

func getImageUrl(tag dom.Tag) (string, bool) {
  rawUrl, ok := tag.Attributes["href"]
  if !ok {
    return "", false
  }

  url := strings.ReplaceAll(rawUrl, "//", "http://")

  return url, true
}

func LoadThreadImages(board, threadID string) (Thread, error) {
  result := Thread{
    Posts: make([]Post, 0),
  }

  threadResult, err := FetchThread(board, threadID)
  if err != nil {
    return result, err
  }

  replies := threadResult.SearchAttribute("class", "fileThumb")
  for _, reply := range replies {
    url, worked := getImageUrl(reply)
    if !worked {
      continue
    }

    post := Post{
      Image: url,
    }

    result.Posts = append(result.Posts, post)
  }

  return result, nil
}
