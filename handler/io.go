package handler

import (
	"bufio"
	"os"
	"strings"
)

func Reader(input *string) string {
	reader := bufio.NewReader(os.Stdin)
	*input, _ = reader.ReadString('\n')
	*input = strings.TrimSpace(*input)
	return *input
}
