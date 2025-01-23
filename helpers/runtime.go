package helpers

import (
	"runtime"
	"strings"
)

// FuncName - returns the name of the function wherein this was invoked.
func FuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	strs := strings.Split((runtime.FuncForPC(pc).Name()), "/")
	return strs[len(strs)-1]
}
