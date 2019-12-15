package dom

import (
  "testing"
)

func TestEqualElements(t *testing.T) {
  tables := []struct{
    InputElement1 Element
    InputElement2 Element
    Output bool
  }{
    {
      InputElement1: Element{
        Type: "div",
        Attributes: map[string]string{},
        Inner: "",
        Children: []*Element{},
      },
      InputElement2: Element{
        Type: "div",
        Attributes: map[string]string{},
        Inner: "",
        Children: []*Element{},
      },
      Output: true,
    },
    {
      InputElement1: Element{
        Type: "div",
        Attributes: map[string]string{},
        Inner: "",
        Children: []*Element{},
      },
      InputElement2: Element{
        Type: "p",
        Attributes: map[string]string{},
        Inner: "",
        Children: []*Element{},
      },
      Output: false,
    },
    {
      InputElement1: Element{
        Type: "div",
        Attributes: map[string]string{
          "class": "test",
        },
        Inner: "",
        Children: []*Element{},
      },
      InputElement2: Element{
        Type: "div",
        Attributes: map[string]string{},
        Inner: "",
        Children: []*Element{},
      },
      Output: false,
    },
    {
      InputElement1: Element{
        Type: "div",
        Attributes: map[string]string{},
        Inner: "test",
        Children: []*Element{},
      },
      InputElement2: Element{
        Type: "div",
        Attributes: map[string]string{},
        Inner: "",
        Children: []*Element{},
      },
      Output: false,
    },
    {
      InputElement1: Element{
        Type: "div",
        Attributes: map[string]string{},
        Inner: "",
        Children: []*Element{
          &Element{
            Type: "div",
            Attributes: map[string]string{},
            Inner: "",
            Children: []*Element{},
          },
        },
      },
      InputElement2: Element{
        Type: "div",
        Attributes: map[string]string{},
        Inner: "",
        Children: []*Element{},
      },
      Output: false,
    },
    {
      InputElement1: Element{
        Type: "div",
        Attributes: map[string]string{},
        Inner: "",
        Children: []*Element{
          &Element{
            Type: "div",
            Attributes: map[string]string{},
            Inner: "",
            Children: []*Element{},
          },
        },
      },
      InputElement2: Element{
        Type: "div",
        Attributes: map[string]string{},
        Inner: "",
        Children: []*Element{
          &Element{
            Type: "p",
            Attributes: map[string]string{},
            Inner: "",
            Children: []*Element{},
          },
        },
      },
      Output: false,
    },
  }

  for _, table := range tables {
    inputElement1 := table.InputElement1
    inputElement2 := table.InputElement2
    output := table.Output

    result := EqualElements(&inputElement1, &inputElement2)

    if result != output {
      t.Errorf("Did not compare Element1 (%+v) and Element1 (%+v) correctly, returned %v, expected %v \n", inputElement1, inputElement2, result, output)
    }
  }
}
