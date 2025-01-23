package formatters

import "github.com/jedib0t/go-pretty/v6/table"

type TableFormattable interface{ TableWriter() table.Writer }

type TableFormatter struct{}

func (f *TableFormatter) Format(data any) ([]byte, error) {
	tw, ok := data.(TableFormattable)
	if !ok {
		return []byte{}, nil
	}

	return []byte(tw.TableWriter().Render()), nil
}
