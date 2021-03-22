package tracex

import (
	"github.com/satori/go.uuid"
	"net/http"

	"github.com/arch3754/tracex/common"
)

type TraceX struct {
	*common.Data
	writer Writer
}

func NewCallerTraceX(metric string, writer Writer) *TraceX {
	d := &common.Data{
		TraceId: uuid.NewV1().String(),
		Metric:  metric,
	}
	d.BuildCaller()
	t := &TraceX{
		Data:   d,
		writer: writer,
	}
	t.writer.Set(d)
	return t
}
func NewRequestTraceX(req *http.Request, metric string, writer Writer) *TraceX {
	d := &common.Data{
		TraceId: uuid.NewV1().String(),
		Metric:  metric,
	}
	d.BuildHTTPRequest(req)
	t := &TraceX{
		Data:   d,
		writer: writer,
	}
	t.writer.Set(d)
	return t
}
func NewTraceX(metric string, writer Writer) *TraceX {

	d := &common.Data{
		TraceId: uuid.NewV1().String(),
		Metric:  metric,
	}
	t := &TraceX{
		Data:   d,
		writer: writer,
	}
	t.writer.Set(d)
	return t
}

func NewRequestAndCallerTraceX(req *http.Request, metric string, writer Writer) *TraceX {
	d := &common.Data{
		TraceId: uuid.NewV1().String(),
		Metric:  metric,
	}
	d.BuildCaller()
	d.BuildHTTPRequest(req)
	t := &TraceX{
		Data:   d,
		writer: writer,
	}
	t.writer.Set(d)
	return t
}

func (t *TraceX) NewChild() *TraceX {
	child := &common.Data{
		User:      t.User,
		Remote:    t.Remote,
		Uri:       t.Uri,
		UserAgent: t.UserAgent,
		Method:    t.Method,
		RequestId: t.RequestId,
		TraceId:   uuid.NewV1().String(),
		Metric:    t.Metric,
		TracePid:  t.TraceId,
	}
	child.BuildCaller()
	n := &TraceX{
		Data:   child,
		writer: t.writer,
	}

	n.writer.Set(child)
	return t
}