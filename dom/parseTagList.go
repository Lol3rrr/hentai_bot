package dom

import (
  "strings"
)

func ParseTagList(tagInput string) (TagList) {
  result := TagList{
    Tags: make([]Tag, 0),
  }

  offset := 0

  startIndex := strings.Index(tagInput, "<")
  for startIndex >= 0 {
    endIndex := strings.Index(tagInput, ">")

    content := tagInput[startIndex + 1: endIndex]
    tagType, tagAtr := parseElementInfo(content)
    tag := Tag{
      Type: tagType,
      Attributes: tagAtr,
    }
    result.Tags = append(result.Tags, tag)

    tagInput = tagInput[endIndex + 1:]
    offset += endIndex + 1
    startIndex = strings.Index(tagInput, "<")
  }

  return result
}
