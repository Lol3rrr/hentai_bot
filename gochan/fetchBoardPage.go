package gochan

import (
  "github.com/Lol3rrr/dom"
)

func FetchBoardPage(board string, page int) (dom.TagList, error) {
  url, err := getPageUrl(board, page)
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
