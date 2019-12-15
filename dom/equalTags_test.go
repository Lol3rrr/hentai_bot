package dom

import (
  "testing"
)

func TestEqualTags(t *testing.T) {
  tables := []struct{
    InputTag1 *Tag
    InputTag2 *Tag
    Output bool
  }{
    {
      InputTag1: &Tag{
        Type: "p",
        Attributes: map[string]string{},
      },
      InputTag2: &Tag{
        Type: "p",
        Attributes: map[string]string{},
      },
      Output: true,
    },
    {
      InputTag1: &Tag{
        Type: "div",
        Attributes: map[string]string{},
      },
      InputTag2: &Tag{
        Type: "p",
        Attributes: map[string]string{},
      },
      Output: false,
    },
    {
      InputTag1: &Tag{
        Type: "p",
        Attributes: map[string]string{
          "id": "yikes",
        },
      },
      InputTag2: &Tag{
        Type: "p",
        Attributes: map[string]string{},
      },
      Output: false,
    },
  }

  for _, table := range tables {
    inputTag1 := table.InputTag1
    inputTag2 := table.InputTag2
    output := table.Output

    result := EqualTags(inputTag1, inputTag2)

    if result != output {
      t.Errorf("Did not compare the Tags (1: %+v ; 2: %+v) correctly, returned %v, expected %v \n", *inputTag1, *inputTag2, result, output)
    }
  }
}
