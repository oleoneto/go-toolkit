package formatters

type SilentFormatter struct{}

func (f SilentFormatter) Format(data any) ([]byte, error) { return []byte{}, nil }
