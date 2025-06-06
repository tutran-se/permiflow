package main

func ClassifyRisk(verbs, resources []string) string {
  if contains(verbs, "*") || contains(resources, "*") {
    return "HIGH"
  }
  if contains(resources, "secrets") {
    return "MEDIUM"
  }
  return "LOW"
}

func contains(list []string, val string) bool {
  for _, v := range list {
    if v == val {
      return true
    }
  }
  return false
}