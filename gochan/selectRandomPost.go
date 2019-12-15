package gochan

func (board *Board) SelectRandomPost() (*Post, bool) {
  if len(board.Threads) < 1 {
    return nil, false
  }
  thread := board.Threads[getRandomNumber(0, len(board.Threads) - 1)]

  if len(thread.Posts) < 1 {
    return nil, false
  }
  post := thread.Posts[getRandomNumber(0, len(thread.Posts) - 1)]

  return &post, true
}
