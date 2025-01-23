package formatters

import "fmt"

type PlainFormatter struct{}

func (PlainFormatter) Format(v interface{}) ([]byte, error) {
	type S interface{ String() string }

	s, ok := v.(S)
	if !ok {
		return []byte(fmt.Sprintf("%+v", v)), nil
	}

	return []byte(s.String()), nil
}
