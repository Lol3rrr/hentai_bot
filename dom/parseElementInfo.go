package dom

import (
  "strings"
)

func parseElementInfo(rawElementInfo string) (string, map[string]string) {
  if len(rawElementInfo) <= 0 {
    return "", map[string]string{}
  }

  typeEnd := strings.Index(rawElementInfo, " ")
  if typeEnd < 0 {
    typeEnd = len(rawElementInfo)
  }

  elemType := rawElementInfo[:typeEnd]
  elemAttributes := getElementParts(rawElementInfo)

  return elemType, elemAttributes
}
