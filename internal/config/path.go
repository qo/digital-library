package config

import (
	"flag"
	"fmt"
	"os"
)

func getPathFromFlag() (string, bool) {
  var path string
  flag.StringVar(&path, "config", "", "path to config")
  flag.Parse()
  
  if path == "" {
    return "", false
  }
  return path, true
}

func getPathFromEnv() (string, bool) {
  path := os.Getenv("DIGITAL_LIBRARY_CONFIG")
  if path == "" {
    return "", false
  }
  return path, true
}

func getPath() (string, error) {
  path, ok := getPathFromFlag()
  if ok {
    return path, nil
  }
  path, ok = getPathFromEnv()
  if ok {
    return path, nil
  }
  return "", fmt.Errorf("config path wasn't specified neither as command-line flag nor as environment variable")
}
