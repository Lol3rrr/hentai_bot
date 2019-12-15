package gochan

import (
  "time"
  "math/rand"
)

func getRandomNumber(min, max int) int {
  rand.Seed(time.Now().UnixNano())
  return (rand.Intn(max - min + 1) + min)
}
