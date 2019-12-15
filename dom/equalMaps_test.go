package dom

import (
  "testing"
)

func TestEqualMaps(t *testing.T) {
  tables := []struct{
    InputMap1 map[string]string
    InputMap2 map[string]string
    Output bool
  }{
    {
      InputMap1: map[string]string{
        "test": "yikes",
      },
      InputMap2: map[string]string{
        "test": "yikes",
      },
      Output: true,
    },
    {
      InputMap1: map[string]string{
        "test": "yikes",
      },
      InputMap2: map[string]string{
      },
      Output: false,
    },
    {
      InputMap1: map[string]string{
        "test": "yikes",
      },
      InputMap2: map[string]string{
        "test2": "yikes",
      },
      Output: false,
    },
    {
      InputMap1: map[string]string{
        "test": "yikes",
      },
      InputMap2: map[string]string{
        "test": "yikes2",
      },
      Output: false,
    },
  }

  for _, table := range tables {
    inputMap1 := table.InputMap1
    inputMap2 := table.InputMap2
    output := table.Output

    result := equalMaps(inputMap1, inputMap2)

    if result != output {
      t.Errorf("Did not compare map1 (%+v) and map2 (%+v) correctly, returned %v, expected %v \n", inputMap1, inputMap2, result, output)
    }
  }
}
