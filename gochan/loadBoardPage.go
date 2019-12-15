package gochan

func LoadBoardPage(board string, page int) ([]Thread, error) {
  result := make([]Thread, 0)

  boardPage, err := FetchBoardPage(board, page)
  if err != nil {
    return result, err
  }

  threadIDs := getPageThreads(boardPage)
  for _, id := range threadIDs {
    thread, err := LoadThreadImages(board, id)
    if err == nil {
      result = append(result, thread)
    }
  }

  return result, nil
}
