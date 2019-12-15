package gochan

import (
  "testing"

  "hentai_bot/dom"
)

func TestGetImageUrl(t *testing.T) {
  tables := []struct{
    InputTag dom.Tag
    ResultString string
    ResultBool bool
  }{
    {
      InputTag: dom.Tag{},
      ResultString: "",
      ResultBool: false,
    },
    {
      InputTag: dom.Tag{
        Attributes: map[string]string{
          "href": "testUrl",
        },
      },
      ResultString: "testUrl",
      ResultBool: true,
    },
    {
      InputTag: dom.Tag{
        Attributes: map[string]string{
          "href": "//testUrl",
        },
      },
      ResultString: "http://testUrl",
      ResultBool: true,
    },
  }

  for _, table := range tables {
    inTag := table.InputTag
    resString := table.ResultString
    resBool := table.ResultBool

    outputString, outputBool := getImageUrl(inTag)

    if outputBool != resBool {
      t.Errorf("")
      continue
    }

    if outputString != resString {
      t.Errorf("")
      continue
    }
  }
}
