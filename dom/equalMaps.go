package dom

func equalMaps(map1, map2 map[string]string) bool {
  if len(map1) != len(map2) {
    return false
  }

  for key := range map1 {
    if _, ok := map2[key]; !ok {
      return false
    }

    if map1[key] != map2[key] {
      return false
    }
  }

  return true
}
