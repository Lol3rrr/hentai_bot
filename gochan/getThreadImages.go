package gochan

import (
  "github.com/Lol3rrr/dom"
)

func getThreadImages(threadResult dom.TagList) []Post {
  result := make([]Post, 0)

  replies := threadResult.SearchAttribute("class", "fileThumb")
  for _, reply := range replies {
    url, worked := getImageUrl(reply)
    if !worked {
      continue
    }

    post := Post{
      Image: url,
    }

    result = append(result, post)
  }

  return result
}
