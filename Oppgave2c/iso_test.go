package test_iso

import (
  "testing"
  "./iso"
)

func TestGreetingExtendedASCII(t *testing.T) {
  iso := iso.GreetingExtendedASCII()
  if !(isExtendedASCII(iso)) {
    t.Fail()
  }
}

func isExtendedASCII(s string) bool {
  for _, c := range s {
    if c > 255 {
      return false
    }
    if c < 127 {
      return false
    }
  }
  return true
}
