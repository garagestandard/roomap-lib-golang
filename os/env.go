package os

import (
	"fmt"
	"os"
)

func Getenv(envname string, def string) (string) {
  value, ok := os.LookupEnv(envname)
  if !ok {
    fmt.Printf("%s not set", envname)
    return def
  }
  return value
}
