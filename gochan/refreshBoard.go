package gochan

import (
  "sync"
)

func (board *Board) Refresh() {
  board.Threads = make([]Thread, 0)

  var waitgroup sync.WaitGroup

  for i := 1; i <= 10; i++ {
    waitgroup.Add(1)

    go loadPage(board, board.Name, i, &waitgroup)
  }

  waitgroup.Wait()
}
