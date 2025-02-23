package parser

import "strings"

func ParseInput(inp string) (string, []string) {
	args := strings.Fields(inp)
	command := args[0]
	text := args[1:]

	return command, text
}
