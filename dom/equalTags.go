package dom

func EqualTags(tag1, tag2 *Tag) bool {
  if tag1.Type != tag2.Type {
    return false
  }
  if !equalMaps(tag1.Attributes, tag2.Attributes) {
    return false
  }

  return true
}
