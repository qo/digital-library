package config

import "os"

func exists(path string) bool {
  _, err := os.Stat(path)
  if os.IsNotExist(err) {
    return false
  }
  return true
}
