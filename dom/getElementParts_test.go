package dom

import (
  "testing"
  "reflect"
)

func TestGetElementParts(t *testing.T) {
  tables := []struct{
    Input string
    Output map[string]string
  }{
    {
      Input: "class=\"test\"",
      Output: map[string]string{
        "class": "test",
      },
    },
    {
      Input: "class=\"test\" id=\"test2\"",
      Output: map[string]string{
        "class": "test",
        "id": "test2",
      },
    },
    {
      Input: "",
      Output: map[string]string{},
    },
  }

  for _, table := range tables {
    input := table.Input
    output := table.Output

    result := getElementParts(input)

    if !reflect.DeepEqual(result, output) {
      t.Errorf("Did not parse the Parts right('%s'), expected '%v' but returned '%v' \n", input, output, result)
      continue
    }
  }
}
