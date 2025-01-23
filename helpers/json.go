package helpers

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strings"
)

func JSONAttributeName(str string) string {
	pattern := regexp.MustCompile(`\.([0-9]+)`)
	scope := strings.Split(str, ": ")[0]
	scope = pattern.ReplaceAllString(scope, "[$1]")

	if scope == "(root)" {
		scope = ""
	}

	if strings.Contains(str, "Additional property") {
		/*
			format:
			- (root): Additional property extra is not allowed
		*/
		p := regexp.MustCompile(`Additional property (.*) is not allowed`)
		name := p.FindStringSubmatch(str)[1]
		return name
	}

	if strings.Contains(str, "required") {
		/*
			format:
				- (root): field_name is required
				- parent.0: field_name is required
		*/
		str = pattern.ReplaceAllString(str, "[$1]")
		name := strings.Split(strings.Trim(strings.SplitAfter(str, ":")[1], " "), " ")
		return strings.TrimPrefix(strings.Join([]string{scope, name[0]}, "."), ".")
	}

	if strings.Contains(str, "Invalid type") {
		/*
			for format:
				- field_name. Expected: typeA, given: typeB
				- field_name.0. Expected: typeA, given: typeB
		*/
		return scope
	}

	return str
}

// JSONPrettyPrint - produces a prettified/beautified JSON string (with proper spacing and indentation).
func JSONPrettyPrint(raw string) string {
	var output bytes.Buffer

	indentPrefix := ""
	indentTokens := "\t"
	err := json.Indent(&output, []byte(raw), indentPrefix, indentTokens)
	if err != nil {
		return raw
	}

	return output.String()
}
