package gochan

import (
  "errors"
)

func getThreadUrl(board, threadID string) (string, error) {
  if len(board) <= 0 {
    return "", errors.New("No Board given")
  }
  if len(threadID) <= 0 {
    return "", errors.New("No Thread ID given")
  }

  return "http://boards.4chan.org/" + board + "/thread/" + threadID, nil
}
