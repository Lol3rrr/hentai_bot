package dom

import (
  "testing"
)

func TestSearchClass(t *testing.T) {
  tables := []struct{
    InputTagList TagList
    InputAttribute string
    InputValue string
    OutputElements []*Tag
  }{
    // Case 1
    {
      InputTagList: TagList{
        Tags: []Tag{
          Tag{
            Type: "p",
            Attributes: map[string]string{
              "class": "test",
            },
          },
          Tag{
            Type: "div",
            Attributes: map[string]string{},
          },
        },
      },
      InputAttribute: "class",
      InputValue: "test",
      OutputElements: []*Tag{
        &Tag{
          Type: "p",
          Attributes: map[string]string{
            "class": "test",
          },
        },
      },
    },
    // Case 2
    {
      InputTagList: TagList{
        Tags: []Tag{},
      },
      InputAttribute: "class",
      InputValue: "test",
      OutputElements: []*Tag{},
    },
    // Case 3
    {
      InputTagList: TagList{
        Tags: []Tag{
          Tag{
            Type: "p",
            Attributes: map[string]string{
              "class": "test1",
            },
          },
        },
      },
      InputAttribute: "class",
      InputValue: "test2",
      OutputElements: []*Tag{},
    },
  }

  for _, table := range tables {
    inputTagList := table.InputTagList
    inputAttribute := table.InputAttribute
    inputValue := table.InputValue
    outputElements := table.OutputElements

    result := inputTagList.SearchAttribute(inputAttribute, inputValue)

    if len(result) != len(outputElements) {
      t.Errorf("Did not search the Tags with Value ('%s') correctly, returned '%+v' expected '%+v' \n", inputValue, result, outputElements)
      continue
    }

    for i := range result {
      if !EqualTags(&result[i], outputElements[i]) {
        t.Errorf("Did not search the Tags with Value ('%s') correctly, returned '%+v' expected '%+v' \n", inputValue, result, outputElements)
        break
      }
    }
  }
}
