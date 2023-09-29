package logger

import (
  "log/slog"
  "os"
  "fmt"
)

const (
  envLocal= "local"
  envDev = "dev"
  envProd = "prod"
)

func Init(env string) (*slog.Logger, error) {
  const errMsg = "can't init logger"

  var logger *slog.Logger

  switch env {
    case envLocal:
      logger = slog.New(
        slog.NewTextHandler(
          os.Stdout, 
          &slog.HandlerOptions{
            Level: slog.LevelDebug,
          },
        ),
      )
    case envDev: 
      logger = slog.New(
        slog.NewJSONHandler(
          os.Stdout, 
          &slog.HandlerOptions{
            Level: slog.LevelDebug,
          },
        ),
      )
    case envProd:
      logger = slog.New(
        slog.NewJSONHandler(
          os.Stdout, 
          &slog.HandlerOptions{
            Level: slog.LevelInfo,
          },
        ),
      )
    default:
      return nil, fmt.Errorf("%s: invalid env specified", errMsg)
  }

  return logger, nil
}
