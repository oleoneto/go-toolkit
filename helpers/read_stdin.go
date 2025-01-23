package helpers

import (
	"io"
	"os"
	"strings"
)

func ReadStdin() ([]byte, error) { return io.ReadAll(os.Stdin) }

func ReadStdinString() (string, error) {
	stdin, err := ReadStdin()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(stdin)), nil
}
