package config

import (
  "io/ioutil"
  "encoding/json"

  "hentai_bot/general"
)

func Load(fileName string) (*Config, error) {
  data, err := ioutil.ReadFile(fileName)
  if general.CheckError(err, "Config") {
    return &Config{}, err
  }

  config := Config{}
  err = json.Unmarshal(data, &config)
  if general.CheckError(err, "Config") {
    return &Config{}, err
  }

  return &config, nil
}
