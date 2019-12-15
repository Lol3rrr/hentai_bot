package gochan

func LoadThreadImages(board, threadID string) (Thread, error) {
  result := Thread{}

  threadResult, err := FetchThread(board, threadID)
  if err != nil {
    return result, err
  }

  result.Posts = getThreadImages(threadResult)

  return result, nil
}
