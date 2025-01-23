package cli

import (
	"sync"
	"time"

	"github.com/drewstinnett/gout/v2"
	"github.com/drewstinnett/gout/v2/config"
	"github.com/drewstinnett/gout/v2/formats/gotemplate"
	gJSON "github.com/drewstinnett/gout/v2/formats/json"
	gYAML "github.com/drewstinnett/gout/v2/formats/yaml"
	"github.com/oleoneto/go-toolkit/cli/formatters"
	"github.com/oleoneto/go-toolkit/files"
	"github.com/spf13/cobra"
)

var once sync.Once
var state *CommandState

func NewCommandState(defaults CommandFlags) *CommandState {
	once.Do(func() {
		state = &CommandState{
			Writer: gout.New(),
			Flags:  defaults,
		}
	})

	return state
}

func (c *CommandState) SetFormatter(cmd *cobra.Command, args []string) {
	switch cmd.Flag("output").Value.String() {
	case "table":
		c.Writer.SetFormatter(&formatters.TableFormatter{})
	case "json":
		c.Writer.SetFormatter(gJSON.Formatter{})
	case "yaml":
		c.Writer.SetFormatter(gYAML.Formatter{})
	case "gotemplate":
		c.Writer.SetFormatter(gotemplate.Formatter{
			Opts: config.FormatterOpts{"template": c.Flags.OutputTemplate},
		})
	case "silent":
		c.Writer.SetFormatter(formatters.SilentFormatter{})
	case "plain":
		c.Writer.SetFormatter(formatters.PlainFormatter{})
	}
}

type CommandState struct {
	Writer             *gout.Gout
	Flags              CommandFlags
	ExecutionStartTime time.Time
	ExecutionExitLog   []any
}

type CommandFlags struct {
	VerboseLogging  bool
	OutputTemplate  string
	OutputFormat    *FlagEnum
	TimeExecutions  bool
	DevelopmentMode bool
	HomeDirectory   string
	CLIConfig       string
	ConfigDir       files.File
	File            FileFlag
	Stdin           string
}
