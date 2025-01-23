package helpers_test

import (
	"fmt"
	"testing"

	"github.com/oleoneto/go-toolkit/helpers"
)

func TestEnumerateArgs(t *testing.T) {
	type args struct {
		counter int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "enumerate with count -1",
			args: args{counter: -1},
			want: "",
		},
		{
			name: "enumerate with count 0",
			args: args{counter: 0},
			want: "",
		},
		{
			name: "enumerate with count 1",
			args: args{counter: 1},
			want: "$1",
		},
		{
			name: "enumerate with count 2",
			args: args{counter: 2},
			want: "$1, $2",
		},
		{
			name: "enumerate with count 3",
			args: args{counter: 3},
			want: "$1, $2, $3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helpers.EnumerateArgs(tt.args.counter, enumerationFunc); got != tt.want {
				t.Errorf("EnumerateArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnumerateArgsOffset(t *testing.T) {
	type args struct {
		counter int
		offset  int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "enumerate with count -1",
			args: args{counter: -1, offset: 1},
			want: "",
		},
		{
			name: "enumerate with count 0",
			args: args{counter: 0, offset: 1},
			want: "",
		},
		{
			name: "enumerate with count 1",
			args: args{counter: 1, offset: 2},
			want: "$3",
		},
		{
			name: "enumerate with count 2",
			args: args{counter: 2, offset: 2},
			want: "$3, $4",
		},
		{
			name: "enumerate with count 3",
			args: args{counter: 3, offset: 3},
			want: "$4, $5, $6",
		},
		{
			name: "enumerate with count 4",
			args: args{counter: 4, offset: len([]string{"a", "b"})},
			want: "$3, $4, $5, $6",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helpers.EnumerateArgsOffset(tt.args.counter, tt.args.offset, enumerationFunc); got != tt.want {
				t.Errorf("EnumerateArgsOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}

var enumerationFunc = func(i, counter int) string { return fmt.Sprintf("$%d", i) }
