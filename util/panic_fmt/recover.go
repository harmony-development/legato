package panic_fmt

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	. "github.com/logrusorgru/aurora"
	"github.com/samber/lo"
)

func RecoverError(clean bool, e interface{}) {
	stack := string(debug.Stack())

	lines := strings.Split(stack, "\n")
	if clean {
		lines = formatStack(lines)
	}

	fmt.Fprintf(os.Stderr, "%s\n", strings.Join(lines, "\n"))
	fmt.Fprintf(os.Stderr, "%s\n", e)
}

func formatStack(lines []string) []string {
	{
		steps := 0
		lines = lo.Map(lines, func(v string, i int) string {
			if strings.Contains(v, "legato") && i%2 == 0 && !strings.Contains(v, "legato/gen") {
				steps++
				highlighted := Red(v).String()
				if steps == 3 {
					highlighted += Yellow(" <- Panic caused here").String()
				}
				return highlighted
			} else {
				return Gray(8, v).String()
			}
		})
	}
	lines = lo.Reverse(lo.Filter(lines, func(_ string, i int) bool { return i != 0 }))
	return lines
}
