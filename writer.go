package tracex

import "github.com/arch3754/tracex/common"

type Writer interface {
	Flush()
	Set(x *common.Data)
}

func (t *TraceX) Flush() {
	t.writer.Flush()
}
