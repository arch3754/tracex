package tracex

import "github.com/arch3754/tracex/common"

type Writer interface {
	Flush() error
	Set(x *common.Data)
}

func (t *TraceX) Flush() error {
	return t.writer.Flush()
}