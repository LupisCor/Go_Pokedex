package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lower_sting := strings.ToLower(text)
	split_string := strings.Fields(lower_sting)
	return split_string
}
