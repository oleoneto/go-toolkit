package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/oleoneto/go-toolkit/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type FileFlag struct{ Data []byte }

func (f FileFlag) String() string { return "" }

func (f FileFlag) Type() string { return "string" }

func (f *FileFlag) Set(value string) error {
	if len(value) < 1 {
		fmt.Println("Error")
		return errors.New("flag not set")
	}

	var err error
	if f.Data, err = os.ReadFile(value); err != nil {
		return err
	}

	return nil
}

func (f *FileFlag) StdinHook(flagName string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if !cmd.Flag(flagName).Changed {
			data, err := helpers.ReadStdin()
			if err != nil {
				os.Exit(1)
			}

			f.Data = data
		}
	}
}

// FileStdinFlag implements pflag.Value
var _ pflag.Value = (*FileFlag)(nil)
