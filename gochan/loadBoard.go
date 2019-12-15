package gochan

import (
  "sync"
)

func LoadBoard(board string) (Board) {
  result := Board{
    Name: board,
    Threads: make([]Thread, 0),
  }

  var waitgroup sync.WaitGroup

  for i := 1; i <= 10; i++ {
    waitgroup.Add(1)

    go loadPage(&result, board, i, &waitgroup)
  }

  waitgroup.Wait()

  return result
}
