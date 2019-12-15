package general

func CheckError(err error, module string) bool {
  if err != nil {
    return true
  }

  return false
}
