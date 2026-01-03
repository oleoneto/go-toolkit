package decoder_test

import (
	"reflect"
	"testing"

	"github.com/oleoneto/go-toolkit/decoder"
)

func Test_StructAttribute_SkipsPastLastChild(t *testing.T) {
	type fields struct {
		value        reflect.Value
		field        reflect.StructField
		parents      []decoder.StructAttribute
		children     []decoder.StructAttribute
		listPosition int
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
				children: []decoder.StructAttribute{},
			},
			want: 0,
		},
		{
			name: "test - 2",
			fields: fields{
				children: []decoder.StructAttribute{{}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := decoder.NewStructAttribute(decoder.NewStructAttributeFields{
				Value:        tt.fields.value,
				Field:        tt.fields.field,
				Parents:      tt.fields.parents,
				Children:     tt.fields.children,
				ListPosition: tt.fields.listPosition,
				IsPrimitive:  tt.fields.isPrimitive,
			})

			if got := sa.SkipsPastLastChild(); got != tt.want {
				t.Errorf("StructAttribute.SkipsPastLastChild() = %v, want %v", got, tt.want)
			}
		})
	}
}
