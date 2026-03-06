package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Scanner = bufio.NewScanner(os.Stdin)

func TanyaUser(pertanyaan string) string {
	fmt.Print(pertanyaan)

	if Scanner.Scan() {
		return strings.TrimSpace(Scanner.Text())
	}

	return ""
}
