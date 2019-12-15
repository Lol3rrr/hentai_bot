package gochan

import (
  "testing"
)

func TestGetThreadUrl(t *testing.T) {
  tables := []struct{
    InputBoard string
    InputThread string
    Output string
    Error bool
  }{
    {
      InputBoard: "b",
      InputThread: "813568638",
      Output: "http://boards.4chan.org/b/thread/813568638",
      Error: false,
    },
    {
      InputBoard: "gif",
      InputThread: "123123",
      Output: "http://boards.4chan.org/gif/thread/123123",
      Error: false,
    },
    {
      InputBoard: "",
      InputThread: "123123",
      Output: "",
      Error: true,
    },
    {
      InputBoard: "gif",
      InputThread: "",
      Output: "",
      Error: true,
    },
  }

  for _, table := range tables {
    inputBoard := table.InputBoard
    inputThread := table.InputThread
    output := table.Output
    error := table.Error

    result, resultErr := getThreadUrl(inputBoard, inputThread)

    if (resultErr != nil) != error {
      t.Errorf("Returned (no) Error '%+v', when it was not expected \n", resultErr)
      continue
    }

    if result != output {
      t.Errorf("Did not create URL correctly, returned '%+v' expected '%+v' \n", result, output)
    }
  }
}
