package gochan

import (
  "testing"
)

func TestGetRandomNumber(t *testing.T) {
  tables := []struct{
    InputMin int
    InputMax int
  }{
    {
      InputMin: 0,
      InputMax: 100,
    },
    {
      InputMin: 0,
      InputMax: 0,
    },
    {
      InputMin: 0,
      InputMax: 1,
    },
  }

  for _, table := range tables {
    inMin := table.InputMin
    inMax := table.InputMax

    output := getRandomNumber(inMin, inMax)

    if output < inMin {
      t.Errorf("The returned Value is lower than the Minimum given, returned '%d', Minimum '%d' \n", output, inMin)
      continue
    }
    if output > inMax {
      t.Errorf("The returned Value is higher than the Maximum given, returned '%d', Maximum '%d' \n", output, inMax)
      continue
    }
  }
}
