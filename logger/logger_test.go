package logger

import (
	"bytes"
	"os"
	"reflect"
	"testing"

	"github.com/drewstinnett/go-output-format/v2/gout"
	"github.com/oleoneto/go-toolkit/types"
)

type (
	S struct{ Value string }
)

func TestNewDefaultLogger(t *testing.T) {
	client.SetFormatter(gout.BuiltInFormatters[PLAINTEXT])

	tests := []struct {
		name string
		want Logger
	}{
		{
			name: "new logger - 1",
			want: Logger{
				paused:      false,
				pendingLogs: types.Queue{},
				format:      PLAINTEXT,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultLogger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogger_Log(t *testing.T) {
	type fields struct {
		paused      bool
		pendingLogs types.Queue
		format      string
	}

	type args struct {
		content  any
		template string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		wantW  string
	}{
		{
			name: "format - plain",
			fields: fields{
				paused:      false,
				pendingLogs: types.Queue{},
				format:      PLAINTEXT,
			},
			args: args{
				content:  S{Value: "Hey!"},
				template: ``,
			},
			wantW: "Hey!",
		},
		{
			name: "format - json",
			fields: fields{
				paused:      false,
				pendingLogs: types.Queue{},
				format:      JSON,
			},
			args: args{
				content:  S{Value: "Hey!"},
				template: ``,
			},
			wantW: `{"Value": "Hey!"}`,
		},
		{
			name: "format - yaml",
			fields: fields{
				paused:      false,
				pendingLogs: types.Queue{},
				format:      YAML,
			},
			args: args{
				content:  S{Value: "Hey!"},
				template: ``,
			},
			wantW: `Value: "Hey!"`,
		},
		{
			name: "format - toml",
			fields: fields{
				paused:      false,
				pendingLogs: types.Queue{},
				format:      TOML,
			},
			args: args{
				content:  S{Value: "Hey!"},
				template: ``,
			},
			wantW: `"Value" = "Hey!"`,
		},
		{
			name: "format - gotemplate",
			fields: fields{
				paused:      false,
				pendingLogs: types.Queue{},
				format:      GOTEMPLATE,
			},
			args: args{
				content:  S{Value: "Hey!"},
				template: `{{.Name}}`,
			},
			wantW: `\n`,
		},
		{
			name: "paused",
			fields: fields{
				paused:      true,
				pendingLogs: types.Queue{},
			},
			args: args{
				content:  S{Value: "Hey!"},
				template: ``,
			},
			wantW: "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			L := Logger{
				paused:      tt.fields.paused,
				pendingLogs: tt.fields.pendingLogs,
				format:      tt.fields.format,
			}
			w := &bytes.Buffer{}
			L.Log(tt.args.content, w, tt.args.template)

			if gotW := w.String(); reflect.DeepEqual(gotW, tt.wantW) {
				t.Errorf("Logger.Log() = %v, want %v.", gotW, tt.wantW)
			}
		})
	}
}

func TestLogger_Pause(t *testing.T) {
	type fields struct {
		delayLogging bool
		pendingLogs  types.Queue
	}

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "pause logger - 1",
			fields: fields{
				delayLogging: false,
				pendingLogs:  types.Queue{},
			},
		},
		{
			name: "pause logger - 2",
			fields: fields{
				delayLogging: true,
				pendingLogs:  types.Queue{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			L := &Logger{
				paused:      tt.fields.delayLogging,
				pendingLogs: tt.fields.pendingLogs,
			}

			L.Pause()

			if !L.paused {
				t.Errorf(`expected logger to be paused`)
			}
		})
	}
}

func TestLogger_Unpause(t *testing.T) {
	type fields struct {
		delayLogging bool
		pendingLogs  types.Queue
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "unpause logger - 1",
			fields: fields{
				delayLogging: false,
				pendingLogs:  types.Queue{},
			},
		},
		{
			name: "unpause logger - 2",
			fields: fields{
				delayLogging: true,
				pendingLogs:  types.Queue{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			L := &Logger{
				paused:      tt.fields.delayLogging,
				pendingLogs: tt.fields.pendingLogs,
			}

			L.Unpause()

			if L.paused {
				t.Errorf(`expected logger to be unpaused`)
			}
		})
	}
}

func TestLogger_Resume(t *testing.T) {
	type fields struct {
		client       gout.Client
		delayLogging bool
	}

	tests := []struct {
		name   string
		fields fields
		logs   []Log
	}{
		{
			name: "resume logger - 1",
			logs: []Log{{content: S{Value: "49"}, writer: os.Stdout}},
			fields: fields{
				client:       gout.Client{Formatter: gout.BuiltInFormatters[PLAINTEXT]},
				delayLogging: false,
			},
		},
		{
			name: "resume logger - 2",
			logs: []Log{
				{content: S{Value: "001"}, writer: os.Stdout},
				{content: S{Value: "002"}, writer: os.Stderr},
				{content: S{Value: "003"}, writer: os.Stdin},
				{content: S{Value: "004"}, writer: os.Stdout},
			},
			fields: fields{
				client:       gout.Client{Formatter: gout.BuiltInFormatters[PLAINTEXT]},
				delayLogging: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			L := &Logger{
				paused:      tt.fields.delayLogging,
				pendingLogs: types.Queue{},
			}

			for _, log := range tt.logs {
				L.save(log.content, log.writer, log.template)
			}

			if L.pendingLogs.IsEmpty() {
				t.Errorf(`expected number of pending logs to be greater than 0, but got %v`, len(tt.logs))
			}

			L.Resume()

			if !L.pendingLogs.IsEmpty() {
				t.Errorf(`expected no more logs, but got %v`, len(tt.logs))
			}
		})
	}
}

func TestLogger_save(t *testing.T) {
	tests := []struct {
		name  string
		logs  []Log
		wantW int
	}{
		{
			name:  "save - 0",
			logs:  []Log{},
			wantW: 0,
		},
		{
			name:  "save - 1",
			logs:  []Log{{content: S{Value: "Invalid1"}, writer: os.Stderr}, {content: S{Value: "Invalid2"}, writer: os.Stderr}},
			wantW: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			L := &Logger{
				paused:      false,
				pendingLogs: types.Queue{},
			}

			for _, log := range tt.logs {
				L.save(log.content, log.writer, log.template)
			}

			if L.NumberOfPendingLogs() != tt.wantW {
				t.Errorf("Logger.save() -> # of pending logs %v, want %v", L.NumberOfPendingLogs(), tt.wantW)
			}
		})
	}
}
