package gochan

import (
  "sync"
)

func loadPage(b *Board, board string, page int, waitgroup *sync.WaitGroup) {
  threads, err := LoadBoardPage(board, page)
  if err != nil {
    return
  }

  b.Threads = append(b.Threads, threads...)

  waitgroup.Done()
}
