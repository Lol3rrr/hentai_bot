package dom

import (
  "testing"
)

func TestParseElementPart(t *testing.T) {
  tables := []struct{
    Input string
    OutputKey string
    OutputValue string
  }{
    {
      Input: "class=\"test\"",
      OutputKey: "class",
      OutputValue: "test",
    },
    {
      Input: "id=\"yikes\"",
      OutputKey: "id",
      OutputValue: "yikes",
    },
    {
      Input: "id=\"\"",
      OutputKey: "id",
      OutputValue: "",
    },
    {
      Input: "id=",
      OutputKey: "",
      OutputValue: "",
    },
    {
      Input: "=\"test\"",
      OutputKey: "",
      OutputValue: "",
    },
  }

  for _, table := range tables {
    input := table.Input
    outputKey := table.OutputKey
    outputValue := table.OutputValue

    resultKey, resultValue := parseElementPart(input)

    if resultKey != outputKey {
      t.Errorf("Did not parse the Key right('%s'), expected '%s' but returned '%s' \n", input, outputKey, resultKey)
      continue
    }
    if resultValue != outputValue {
      t.Errorf("Did not parse the Value right('%s'), expected '%s' but returned '%s' \n", input, outputValue, resultValue)
      continue
    }
  }
}
