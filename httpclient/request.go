package httpclient

import (
	"bytes"
	"io"
	"net/http"
)

type Body struct{ state bytes.Buffer }

func (b *Body) Close() error { return io.NopCloser(&b.state).Close() }

func (b *Body) Read(buffer []byte) (int, error) { return b.state.Read(buffer) }

func NewBody(data []byte) io.ReadCloser {
	if data == nil || len(data) < 1 {
		return http.NoBody
	}
	return &Body{*bytes.NewBuffer(data)}
}

// Body implements io.ReadCloser
var _ io.ReadCloser = (*Body)(nil)
