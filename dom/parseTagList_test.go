package dom

import (
  "testing"
)

func TestParseTagList(t *testing.T) {
  tables := []struct{
    InputDom string
    Output TagList
  }{
    {
      InputDom: "<p>",
      Output: TagList{
        Tags: []Tag{
          Tag{
            Type: "p",
            Attributes: map[string]string{},
          },
        },
      },
    },
    {
      InputDom: "<p class=\"yikes\">",
      Output: TagList{
        Tags: []Tag{
          Tag{
            Type: "p",
            Attributes: map[string]string{
              "class": "yikes",
            },
          },
        },
      },
    },
  }

  for _, table := range tables {
    inputDom := table.InputDom
    output := table.Output

    result := ParseTagList(inputDom)

    if len(result.Tags) != len(output.Tags) {
      t.Errorf("Did not parse the Tag Lists correctly \n")
      continue
    }
  }
}
