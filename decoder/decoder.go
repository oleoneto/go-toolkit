package decoder

import (
	"regexp"

	"github.com/invopop/jsonschema"
	"github.com/oleoneto/go-toolkit/helpers"
	"github.com/xeipuuv/gojsonschema"
)

// Replacement for the standard `json.Unmarshal` implementation.
// It deserializes a JSON object into a Go struct. This function does not
// Panic when the value for a JSON field is incompatible with the type set in the struct.
//
// You're allowed to pass a list of `SchemaValidationType` to use while deserializing the JSON payload
// into your struct. They are:
//   - `ATTRIBUTE_MUST_BE_PRESENT`:
//
// checks if a required field is absent from the JSON payload.
//   - `ADDITIONAL_PROPERTY`:
//
// checks if an unknown field was passed in the JSON payload.
//   - `INVALID_TYPE`:
//
// checks if the type of a JSON attribute in the payload is compatible with the underlying type of the Go struct field.
//
// Usage:
//
//	type User struct {
//		Id 	   int 		`json:"id"`
//		Name   string 	`json:"name" validate:"is_present"`
//		Emails []string	`json:"emails,omitempty"`
//	}
//
//	payload := []byte(`{"name": 42, "emails": ["test@example.com", 0]}`)
//	parsedValues, errs := Decode(payload, User{}, options)
//	/*
//	expected errors:
//	[
//		"id - REQUIRED_ATTRIBUTE_MISSING",
//		"name - INVALID_DATA_TYPE",
//		"emails[1] - INVALID_DATA_TYPE"
//	]
//	*/
func Decode(data []byte, model any, options DecoderOptions) map[string][]string {
	validations := make(map[string][]string, 0)

	if options.BeforeHook != nil {
		data = options.BeforeHook(data, model)
	}

	SetValuesFromBytes(model, data)

	afterFunc := func(validations map[string][]string) map[string][]string {
		return validations
	}

	if options.AfterHook != nil {
		afterFunc = options.AfterHook
	}

	if len(data) == 0 || len(options.Rules) == 0 {
		return afterFunc(validations)
	}

	reflector := new(jsonschema.Reflector)
	reflector.RequiredFromJSONSchemaTags = true
	reflector.AllowAdditionalProperties = !helpers.Contains(options.Rules, ADDITIONAL_PROPERTY)

	schema := reflector.Reflect(model)
	for _, t := range options.JSONOverrides {
		if _, ok := schema.Definitions[t.GoType]; ok {
			schema.Definitions[t.GoType].Type = t.JSONType
		}
	}

	decoded, _ := schema.MarshalJSON()

	result, verr := gojsonschema.Validate(
		gojsonschema.NewBytesLoader(decoded),
		gojsonschema.NewBytesLoader(data),
	)

	if verr != nil {
		validations["_"] = []string{DecodingErrors["invalid_payload"]}
		return afterFunc(validations)
	}

	res := helpers.Filter(result.Errors(), func(index int, err gojsonschema.ResultError) bool {
		return helpers.Contains(options.Rules, SchemaValidationRule(err.Type()))
	})

	for _, err := range res {
		name := helpers.JSONAttributeName(err.String())
		normalizedName := regexp.MustCompile(`\[\d+\]`).ReplaceAllString(name, "")
		validations[normalizedName] = []string{DecodingErrors[err.Type()]}
	}

	return afterFunc(validations)
}

type (
	SchemaValidationRule string

	JSONTypeOverride struct {
		// A string representing the Go struct type.
		// For example: Google's uuid.UUID type would be UUID.
		GoType string

		// The JSON representation of the Go type.
		// For example: number, string, object, array.
		JSONType string
	}

	DecoderOptions struct {
		// Set of rules that should be checked when validation the provided data against the Go struct.
		Rules []SchemaValidationRule

		// Set of Go types whose JSON representation you which to manually override.
		// By default, Go structs are treated as JSON objects. However, you may have a
		// custom type whose JSON representation may simply be a number or a string.
		JSONOverrides []JSONTypeOverride

		// A function that runs before the decoder starts processing the data.
		// This could be used for setting/unsetting values in the provided bytes array.
		BeforeHook func(data []byte, model any) []byte

		// A function that runs after the decoder is done processing the data.
		// This could be used for ignoring certain errors or providing custom error messages.
		AfterHook func(validations map[string][]string) map[string][]string
	}
)

const (
	ADDITIONAL_PROPERTY SchemaValidationRule = "additional_property_not_allowed"
	REQUIRED_ATTRIBUTE  SchemaValidationRule = "required"
	INVALID_TYPE        SchemaValidationRule = "invalid_type"
)

var DecodingErrors = map[string]string{
	"required":                        "REQUIRED_ATTRIBUTE_MISSING",
	"invalid_payload":                 "INVALID_PAYLOAD",
	"invalid_type":                    "INVALID_TYPE",
	"additional_property_not_allowed": "ADDITIONAL_PROPERTY",
}
