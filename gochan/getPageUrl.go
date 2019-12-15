package gochan

import (
  "errors"
  "strconv"
)

func getPageUrl(board string, page int) (string, error) {
  if page < 1 {
    return "", errors.New("Page has to start at 1")
  }
  if len(board) <= 0 {
    return "", errors.New("No Board given")
  }

  url := "http://boards.4chan.org/" + board + "/"

  if page > 1 {
    pageStr := strconv.Itoa(page)
    url = url + pageStr
  }

  return url, nil
}
