package dom

import (
  "strings"
)

func parseElementPart(input string) (string, string) {
  divider := "=\""
  dividerIndex := strings.Index(input, divider)
  if dividerIndex == -1 || dividerIndex == 0 {
    return "", ""
  }

  key := input[:dividerIndex]
  value := input[dividerIndex + len(divider):len(input) - 1]

  return key, value
}
