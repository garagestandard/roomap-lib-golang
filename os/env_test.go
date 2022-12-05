package os

import (
	"os"
	"testing"
)

func TestGetenv(t *testing.T) {
  // default value is set
  envName := "RADIOHEAD"
  result := Getenv(envName, "kid a")
  expect := "kid a"
  if result != expect {
    t.Error("result:", result, " expect:", expect)
  }

  // default value is ""
  result = Getenv(envName, "")
  expect = ""
  if result != expect {
    t.Error("result:", result, " expect:", expect)
  }

  // env var is ""
  result = Getenv("", "")
  expect = ""
  if result != expect {
    t.Error("result:", result, " expect:", expect)
  }

  // env var is set
  envValue := "OK Computer"
  os.Setenv(envName, envValue)
  result = Getenv(envName, "")
  expect = envValue
  if result != expect {
    t.Error("result:", result, " expect:", expect)
  }
 

}
