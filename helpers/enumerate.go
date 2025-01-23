package helpers

import (
	"fmt"
	"strings"
)

// EnumerateArgs - returns an SQL-ready string representation of a counter.
//
// As a special case, a counter <= 0 returns an empty string.
//
// Usage examples:
//
//	EnumerateArgs(1, func(index, _ int) string{ return fmt.Sprintf("$%d", index) }) // $1, $2, $3
//	EnumerateArgs(len([]string{"a", "b"}), func(index, _ int) string{ return fmt.Sprintf("$%d", index) }) // $1, $2
func EnumerateArgs(counter int, valueFunc func(index, counter int) string) string {
	return EnumerateArgsOffset(counter, 0, valueFunc)
}

// EnumerateArgsOffset - returns an SQL-ready string representation of a counter.
//
// As a special case, a counter <= 0 returns an empty string.
//
// Usage:
//
//	EnumerateArgsOffset(3, 0, func(i, c int) string { return "?" }) // ?, ?, ?
//	EnumerateArgsOffset(3, 0, func(i, c int) string { return fmt.Sprintf("$%d", i) }) // $1, $2, $3
func EnumerateArgsOffset(counter, offset int, valueFunc func(index, counter int) string) string {
	var out string

	for i := 1; i <= counter; i++ {
		out += fmt.Sprintf(`%v, `, valueFunc(i+offset, counter))
	}

	out = strings.TrimSpace(out)
	out = strings.TrimRight(out, ",")

	return out
}
