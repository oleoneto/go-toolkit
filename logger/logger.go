package logger

import (
	"context"
	"io"

	"github.com/drewstinnett/go-output-format/v2/formats/gotemplate"
	"github.com/drewstinnett/go-output-format/v2/gout"
	"github.com/oleoneto/go-toolkit/types"
)

type (
	Formattable interface {
		String() string
	}

	Log struct {
		content  any
		writer   io.Writer
		template string
	}

	Logger struct {
		pendingLogs types.Queue
		paused      bool
		format      string
	}

	LoggerOptions struct {
		Format       string
		DelayLogging bool
	}
)

const (
	CSV        string = "csv"
	GOTEMPLATE string = "gotemplate"
	JSON       string = "json"
	PLAINTEXT  string = "plain"
	TOML       string = "toml"
	XML        string = "xml"
	YAML       string = "yaml"
)

var client, _ = gout.New()

func init() {
	client.SetFormatter(gout.BuiltInFormatters[PLAINTEXT])
}

func NewLogger(options LoggerOptions) Logger {
	client.SetFormatter(gout.BuiltInFormatters[options.Format])

	return Logger{
		paused:      options.DelayLogging,
		format:      options.Format,
		pendingLogs: types.NewQueue(),
	}
}

func NewDefaultLogger() Logger {
	return NewLogger(LoggerOptions{
		DelayLogging: false,
		Format:       PLAINTEXT,
	})
}

func (L *Logger) Log(content any, w io.Writer, template string) {
	if L.paused {
		L.save(content, w, template)
		return
	}

	if L.format == GOTEMPLATE && template != "" {
		ctx := context.WithValue(context.Background(), gotemplate.TemplateField{}, template)
		b, _ := client.Formatter.FormatWithContext(ctx, content)
		client.Print(string(b))
		return
	}

	client.SetWriter(w)
	client.Print(content)
}

func (L *Logger) Pause() {
	L.paused = true
}

func (L *Logger) Unpause() {
	L.paused = false
}

func (L *Logger) Resume() {
	L.Unpause()

	if data := L.pendingLogs.Dequeue(); data != nil {
		for L.pendingLogs.Size() > 0 {
			if log, ok := data.(Log); ok {
				L.Log(log.content, log.writer, log.template)
			}

			data = L.pendingLogs.Dequeue()
		}
	}
}

func (L *Logger) NumberOfPendingLogs() int {
	return L.pendingLogs.Size()
}

func (L *Logger) save(content any, w io.Writer, template string) {
	L.pendingLogs.Enqueue(Log{content: content, writer: w, template: template})
}
