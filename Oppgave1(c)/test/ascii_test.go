package test_ascii

import (
  "testing"
  "./ascii"
)

func TestGreetingASCII(t *testing.T) {
  ascii := ascii.GreetingASCII()
  if !(isASCII(ascii)) {
    t.Fail()
  }
}

func isASCII(s string) bool {
  for _, c := range s {
    if c > 127 {
      return false
    }
  }
  return true
}
