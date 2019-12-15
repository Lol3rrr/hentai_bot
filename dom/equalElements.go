package dom

func EqualElements(elem1, elem2 *Element) bool {
  if elem1.Type != elem2.Type {
    return false
  }
  if !equalMaps(elem1.Attributes, elem2.Attributes) {
    return false
  }
  if elem1.Inner != elem2.Inner {
    return false
  }

  if len(elem1.Children) != len(elem2.Children) {
    return false
  }

  for i := range elem1.Children {
    if !EqualElements(elem1.Children[i], elem2.Children[i]) {
      return false
    }
  }

  return true
}
