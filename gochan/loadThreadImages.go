package gochan

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
