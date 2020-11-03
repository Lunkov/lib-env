package env

import (
  "testing"
)

func TestGetDefault(t *testing.T) {
  expect := "redis://localhost2/"
  res := Get("REDIS_URL_123", expect)
  if res != expect {
    t.Error(
      "For", "Environment Get",
      "expected", expect,
      "got", res,
    )
  }
}

func TestGetIntDefault(t *testing.T) {
  expect := 10
  res := GetInt("TIMEOUT_123", expect)
  if res != expect {
    t.Error(
      "For", "Environment Get Integer",
      "expected", expect,
      "got", res,
    )
  }
}

func TestWaitFile(t *testing.T) {
  
  res := WaitFile("env.go", 1)
  if res != true {
    t.Error(
      "For", "Environment WaitFile",
      "expected", true,
      "got", res,
    )
  }

  res = WaitFile("env.go111", 1)
  if res != false {
    t.Error(
      "For", "Environment WaitFile",
      "expected", false,
      "got", res,
    )
  }

}

