package gochan

func LoadBoardPage(board string, page int) ([]Thread, error) {
  result := make([]Thread, 0)

  boardPage, err := FetchBoardPage(board, page)
  if err != nil {
    return result, err
  }
  threads := boardPage.SearchAttribute("class", "thread")
  for _, thread := range threads {
    rawID := thread.Attributes["id"]
    id := rawID[1:]

    thread, err := LoadThreadImages(board, id)
    if err != nil {
      continue
    }

    result = append(result, thread)
  }

  return result, nil
}
