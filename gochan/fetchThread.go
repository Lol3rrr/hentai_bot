package gochan

import (
  "github.com/Lol3rrr/dom"
)

func FetchThread(board, threadID string) (dom.TagList, error) {
  url, err := getThreadUrl(board, threadID)
  if err != nil {
    return dom.TagList{}, err
  }

  response, err := MakeBasicGetRequest(url)
  if err != nil {
    return dom.TagList{}, err
  }

  result := dom.ParseTagList(response)

  return result, nil
}
