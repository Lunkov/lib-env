package env

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestGetDefault(t *testing.T) {
  expect := "redis://localhost2/"
  res := Get("REDIS_URL_123", expect)

  assert.Equal(t, expect, res)
}

func TestGetIntDefault(t *testing.T) {
  expect := 10
  res := GetInt("TIMEOUT_123", expect)
  
  assert.Equal(t, expect, res)
}

func TestWaitFile(t *testing.T) {
  
  res := WaitFile("env.go", 1)
  assert.Equal(t, true, res)

  res = WaitFile("env.go111", 1)
  assert.Equal(t, false, res)
}

