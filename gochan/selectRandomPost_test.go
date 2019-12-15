package gochan

import (
  "testing"
)

func TestSelectRandomPost(t *testing.T) {
  tables := []struct{
    InputBoard Board
    ResultBool bool
  }{
    {
      InputBoard: Board{},
      ResultBool: false,
    },
    {
      InputBoard: Board{
        Threads: []Thread{
          {
            Posts: []Post{},
          },
        },
      },
      ResultBool: false,
    },
    {
      InputBoard: Board{
        Threads: []Thread{
          {
            Posts: []Post{
              {
                Image: "imageTest",
                Content: "contentTest",
              },
            },
          },
        },
      },
      ResultBool: true,
    },
  }

  for _, table := range tables {
    inBoard := table.InputBoard
    resBool := table.ResultBool

    outputPost, outputBool := inBoard.SelectRandomPost()

    if outputBool != resBool {
      t.Errorf("Did not return the correct Bool value, returned '%v', expected '%v' \n", outputBool, resBool)
      continue
    }

    if outputPost == nil && resBool {
      t.Errorf("Returned a nil post when it should have returned a valid post")
      continue
    }
  }
}
