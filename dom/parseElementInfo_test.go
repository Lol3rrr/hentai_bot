package dom

import (
  "testing"
  "reflect"
)

func TestParseElementInfo(t *testing.T) {
  tables := []struct{
    Input string
    OutputType string
    OutputAttributes map[string]string
  }{
    {
      Input: "p",
      OutputType: "p",
      OutputAttributes: map[string]string{},
    },
    {
      Input: "p class=\"test\"",
      OutputType: "p",
      OutputAttributes: map[string]string{
        "class": "test",
      },
    },
    {
      Input: "p id=\"test\"",
      OutputType: "p",
      OutputAttributes: map[string]string{
        "id": "test",
      },
    },
    {
      Input: "p id=\"test\" class=\"test\"",
      OutputType: "p",
      OutputAttributes: map[string]string{
        "id": "test",
        "class": "test",
      },
    },
    {
      Input: "",
      OutputType: "",
      OutputAttributes: map[string]string{},
    },
  }

  for _, table := range tables {
    input := table.Input
    outputType := table.OutputType
    outputAttributes := table.OutputAttributes

    resultType, resultAttributes := parseElementInfo(input)

    if resultType != outputType {
      t.Errorf("Did not parse the Type right('%s'), expected '%s' but returned '%s' \n", input, outputType, resultType)
      continue
    }
    if !reflect.DeepEqual(resultAttributes, outputAttributes) {
      t.Errorf("Did not parse the Attributes right('%s'), expected '%s' but returned '%s' \n", input, outputAttributes, resultAttributes)
      continue
    }
  }
}
