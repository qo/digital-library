package main

import (
	"fmt"
	"os"

	"github.com/qo/digital-library/internal/config"
)

func main() {
  cfg, err := config.Load()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  fmt.Println(cfg)
} 
