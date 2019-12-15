package general

import (
  "errors"
  "testing"
)

func TestCheckError(t *testing.T) {
  tables := []struct{
    Input error
    InputModule string
    Output bool
  }{
    {
      Input: errors.New("Test"),
      InputModule: "Testing",
      Output: true,
    },
    {
      Input: nil,
      InputModule: "Testing",
      Output: false,
    },
  }

  for _, table := range tables {
    input := table.Input
    inputModule := table.InputModule
    output := table.Output

    result := CheckError(input, inputModule)

    if result != output {
      t.Errorf("Did not check for Error correctly, returned '%+v' expected '%+v' \n", result, output)
    }
  }
}
