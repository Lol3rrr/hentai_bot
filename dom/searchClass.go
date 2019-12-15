package dom

func (tagList *TagList) SearchAttribute(attribute, value string) ([]Tag) {
  result := make([]Tag, 0)

  for _, tmpTag := range tagList.Tags {
    elemValue, ok := tmpTag.Attributes[attribute]
    if !ok {
      continue
    }
    if elemValue != value {
      continue
    }

    result = append(result, tmpTag)
  }

  return result
}
