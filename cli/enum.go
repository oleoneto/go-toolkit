package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

type FlagEnum struct {
	Allowed []string
	Default string
}

func (e FlagEnum) String() string { return e.Default }

func (e *FlagEnum) Type() string { return "string" }

func (e *FlagEnum) Set(value string) error {
	isIncluded := func(opts []string, v string) bool {
		for _, opt := range opts {
			if v == opt {
				return true
			}
		}

		return false
	}

	if !isIncluded(e.Allowed, value) {
		return fmt.Errorf("%s is not a supported output format: %s", value, strings.Join(e.Allowed, ","))
	}

	e.Default = value
	return nil
}

// FileEnum implements pflag.Value
var _ pflag.Value = (*FlagEnum)(nil)
