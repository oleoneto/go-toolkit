package decoder_test

import (
	"reflect"
	"testing"

	"github.com/oleoneto/go-toolkit/decoder"
)

func TestStructAttribute_SkipsPastLastChild(t *testing.T) {
	type fields struct {
		Value        reflect.Value
		Field        reflect.StructField
		Parents      []decoder.StructAttribute
		Children     []decoder.StructAttribute
		ListPosition int
		isPrimitive  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "test - 1",
			fields: fields{
				Children: []decoder.StructAttribute{},
			},
			want: 0,
		},
		{
			name: "test - 2",
			fields: fields{
				Children: []decoder.StructAttribute{{}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := decoder.NewStructAttribute(decoder.NewStructAttributeFields{
				Value:        tt.fields.Value,
				Field:        tt.fields.Field,
				Parents:      tt.fields.Parents,
				Children:     tt.fields.Children,
				ListPosition: tt.fields.ListPosition,
				IsPrimitive:  tt.fields.isPrimitive,
			})

			if got := sa.SkipsPastLastChild(); got != tt.want {
				t.Errorf("StructAttribute.SkipsPastLastChild() = %v, want %v", got, tt.want)
			}
		})
	}
}
