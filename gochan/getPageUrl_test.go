package gochan

import (
  "testing"
)

func TestGetPageUrl(t *testing.T) {
  tables := []struct{
    InputBoard string
    InputPage int
    Output string
    Error bool
  }{
    {
      InputBoard: "b",
      InputPage: 1,
      Output: "http://boards.4chan.org/b/",
      Error: false,
    },
    {
      InputBoard: "b",
      InputPage: 2,
      Output: "http://boards.4chan.org/b/2",
      Error: false,
    },
    {
      InputBoard: "b",
      InputPage: 0,
      Output: "",
      Error: true,
    },
    {
      InputBoard: "",
      InputPage: 1,
      Output: "",
      Error: true,
    },
  }

  for _, table := range tables {
    inputBoard := table.InputBoard
    inputPage := table.InputPage
    output := table.Output
    error := table.Error

    result, resultErr := getPageUrl(inputBoard, inputPage)

    if (resultErr != nil) != error {
      t.Errorf("Returned (no) Error '%+v', when it was not expected \n", resultErr)
      continue
    }

    if result != output {
      t.Errorf("Did not create URL correctly, returned '%+v' expected '%+v' \n", result, output)
    }
  }
}
