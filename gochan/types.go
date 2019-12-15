package gochan

type Board struct {
  Name string
  Threads []Thread
}

type Thread struct {
  Posts []Post
}

type Post struct {
  Image string
  Content string
}
