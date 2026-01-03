package decoder

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/oleoneto/go-toolkit/helpers"
)

// This type abstracts both `reflect.Value` and `reflect.StructField` types.
type StructAttribute struct {
	Value        reflect.Value
	Field        reflect.StructField
	Children     []StructAttribute
	ListPosition int
	isPrimitive  bool

	name *string
}

type NewStructAttributeFields struct {
	Value             reflect.Value
	Field             reflect.StructField
	Parents, Children []StructAttribute
	ListPosition      int
	IsPrimitive       bool
}

type StructAttributes []StructAttribute

// Sets and returns the name of the field properly scoped under its parents.
//
// Usage:
//
// Imagine you have the following StructAttribute:
//
//	sa := StructAttribute{
//		Parents: []StructAttribute{parentA, listB}
//		Field: reflect.StructField{
//			Name:    "attribute_name",
//			...
//		}
//	}
//
//	sa.assignName(parents...) // -> "parentA.listB[i].attribute_name"
func (sa *StructAttribute) assignName(parents ...StructAttribute) string {
	if len(parents) < 1 {
		sa.name = helpers.PointerTo(GetJSONTagValue(sa.Field))
		return *sa.name
	}

	scope := parents[len(parents)-1].FullName()

	// Adds the array notation to the slice/array field
	if sa.ListPosition >= 0 {
		scope = strings.Join([]string{scope, fmt.Sprint("[", sa.ListPosition, "]")}, "")
	}

	if sa.isPrimitive {
		sa.name = helpers.PointerTo(scope)
		return *sa.name
	}

	fullName := strings.Join([]string{scope, GetJSONTagValue(sa.Field)}, ".")

	// Ensures field name is never prefixed or suffixed by a dot (.)
	sa.name = helpers.PointerTo(strings.TrimSuffix(strings.TrimPrefix(fullName, "."), "."))

	return *sa.name
}

func (sa *StructAttribute) FullName() string { return *sa.name }

func (sa *StructAttribute) SkipsPastLastChild() int {
	if len(sa.Children) == 0 {
		return 0
	}

	n := 1 + len(sa.Children)
	for _, child := range sa.Children {
		n += 1 + child.SkipsPastLastChild()
	}

	return n
}

func NewStructAttribute(args NewStructAttributeFields) *StructAttribute {
	sa := &StructAttribute{
		Value:        args.Value,
		Field:        args.Field,
		Children:     args.Children,
		ListPosition: args.ListPosition,
		isPrimitive:  args.IsPrimitive,
	}

	sa.assignName(args.Parents...)

	return sa
}
