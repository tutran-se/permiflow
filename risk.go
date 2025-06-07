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
