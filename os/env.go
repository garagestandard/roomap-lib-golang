package os

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
  godotenv.Load()
}

func Getenv(envname string, def string) (string) {
  value, ok := os.LookupEnv(envname)
  if !ok {
    fmt.Printf("%s not set\n", envname)
    return def
  }
  return value
}
