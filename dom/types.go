package dom

type Element struct {
  Type string
  Attributes map[string]string
  Inner string
  Children []*Element
}

type TagList struct {
  Tags []Tag
}

type Tag struct {
  Type string
  Attributes map[string]string
}
