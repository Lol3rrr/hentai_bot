package gochan

import (
  "reflect"
  "testing"

  "hentai_bot/dom"
)

func TestGetPageThreads(t *testing.T) {
  tables := []struct{
    InputBoardPage dom.TagList
    ResultIDs []string
  }{
    {
      InputBoardPage: dom.TagList{},
      ResultIDs: []string{},
    },
    {
      InputBoardPage: dom.TagList{
        Tags: []dom.Tag{
          {
            Attributes: map[string]string{
              "class": "thread",
            },
          },
        },
      },
      ResultIDs: []string{},
    },
    {
      InputBoardPage: dom.TagList{
        Tags: []dom.Tag{
          {
            Attributes: map[string]string{
              "class": "thread",
              "id": "tidTest",
            },
          },
        },
      },
      ResultIDs: []string{
        "idTest",
      },
    },
  }

  for _, table := range tables {
    inBoardPage := table.InputBoardPage
    resIDs := table.ResultIDs

    outputIDs := getPageThreads(inBoardPage)

    if !reflect.DeepEqual(outputIDs, resIDs) {
      t.Errorf("Did not return the IDs correctly, returned '%v', expected '%v' \n", outputIDs, resIDs)
      continue
    }
  }
}
