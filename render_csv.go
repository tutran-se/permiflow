package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func WriteCSV(bindings []AccessBinding, filename string) {
  f, _ := os.Create(filename)
  defer f.Close()
  writer := csv.NewWriter(f)
  defer writer.Flush()

  writer.Write([]string{"Subject", "Namespace", "Role", "Verbs", "Resources", "RiskLevel"})
  for _, b := range bindings {
    writer.Write([]string{
      b.Subject,
      b.Namespace,
      b.Role,
      fmt.Sprint(b.Verbs),
      fmt.Sprint(b.Resources),
      b.RiskLevel,
    })
  }
}