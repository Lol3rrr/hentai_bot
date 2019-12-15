package gochan

import (
  "net/http"
  "io/ioutil"
)

func MakeBasicGetRequest(url string) (string, error) {
  client := &http.Client{}
  req, _ := http.NewRequest("GET", url, nil)
  resp, err := client.Do(req)
  if err != nil {
    return "", err
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  return string(body), nil
}
