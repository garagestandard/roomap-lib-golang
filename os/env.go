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
  if envname == "" {
    fmt.Printf("Getenv: envname is not se\n")
    return "";
  }
  value, ok := os.LookupEnv(envname)
  if !ok {
    fmt.Printf("Getenv: %s not set\n", envname)
    if def != "" {
      fmt.Printf("Getenv: return default value:%s\n", def)
      return def
    }
    return ""
  }
  return value
}
