package server

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	. "github.com/logrusorgru/aurora"
	"github.com/samber/lo"
)

func Recover(clean bool) fiber.Handler {
	return recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			stack := string(debug.Stack())

			lines := strings.Split(stack, "\n")
			if clean {
				lines = filterStack(lines)
			}

			fmt.Fprintf(os.Stderr, "%s\n", strings.Join(lines, "\n"))
			fmt.Fprintf(os.Stderr, "%s\n", e)

			c.Status(500).SendString(fmt.Sprintf("%v", e))
		},
	})
}

func filterStack(lines []string) []string {
	lines = lo.Map(lines, func(v string, i int) string {
		if strings.Contains(v, "legato") && i%2 == 0 && !strings.Contains(v, "legato/gen") {
			return Red(v).String()
		} else {
			return Gray(8, v).String()
		}
	})
	lines = lo.Reverse(lo.Filter(lines, func(_ string, i int) bool { return i != 0 }))
	lines = lo.Map(lines, func(s string, i int) string { return strings.TrimPrefix(s, "\t") })
	return lines
}
