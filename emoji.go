package main

import (
	"os"
)

func emoji(s string) string {
  if os.Getenv("PERMIFLOW_NO_EMOJI") == "true" {
    return ""
  }
  return s
}
