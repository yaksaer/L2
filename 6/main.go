package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type flags struct {
	f *int
	d *string
	s *bool
}

func newFlags() *flags {
	flags := new(flags)
	flags.f = flag.Int("f", 1, "selected fields")
	flags.d = flag.String("d", "\t", "delimiter")
	flags.s = flag.Bool("s", false, "separated")
	return flags
}

func cut(text string, flags flags) string {

	if *flags.s {
		if !strings.Contains(text, *flags.d) {
			return ""
		}
	}
	res := strings.Split(text, *flags.d)

	if *flags.f <= len(res) {
		var sb strings.Builder
		sb.WriteString(res[*flags.f-1])
		sb.WriteString("\n")
		return sb.String()
	}

	return ""

}

func main() {
	flags := newFlags()
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	text := make([]string, 0, 3)
	run := true
	for run {
		scanner.Scan()
		line := scanner.Text()
		switch line {
		case "\\EOF":
			run = false
		default:
			text = append(text, line)
		}
	}

	for _, v := range text {
		fmt.Print(cut(v, *flags))
	}
}
