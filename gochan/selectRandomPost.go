package gochan

func (board *Board) SelectRandomPost() (*Post, bool) {
  thread := board.Threads[getRandomNumber(0, len(board.Threads))]

  post := thread.Posts[getRandomNumber(0, len(thread.Posts))]

  return &post, true
}
