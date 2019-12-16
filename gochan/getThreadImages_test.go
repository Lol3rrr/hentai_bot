package gochan

import (
  "reflect"
  "testing"

  "github.com/Lol3rrr/dom"
)

func TestGetThreadImages(t *testing.T) {
  tables := []struct{
    InputThread dom.TagList
    ResultPosts []Post
  }{
    {
      InputThread: dom.TagList{},
      ResultPosts: []Post{},
    },
    {
      InputThread: dom.TagList{
        Tags: []dom.Tag{
          {
            Attributes: map[string]string{
              "class": "fileThumb",
            },
          },
        },
      },
      ResultPosts: []Post{},
    },
    {
      InputThread: dom.TagList{
        Tags: []dom.Tag{
          {
            Attributes: map[string]string{
              "class": "fileThumb",
              "href": "urlTest",
            },
          },
        },
      },
      ResultPosts: []Post{
        {
          Image: "urlTest",
        },
      },
    },
  }

  for _, table := range tables {
    inThread := table.InputThread
    resPosts := table.ResultPosts

    outputPosts := getThreadImages(inThread)

    if !reflect.DeepEqual(outputPosts, resPosts) {
      t.Errorf("Did not get the Images correctly, returned '%v', expected '%v' \n", outputPosts, resPosts)
      continue
    }
  }
}
