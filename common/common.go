package common

import (
	"encoding/json"
	"net/http"
	"runtime"
)

type Data struct {
	FuncName  string `json:"func_name,omitempty" xorm:"func_name" db:"func_name"`
	Line      int    `json:"line,omitempty" xorm:"line" db:"line"`
	File      string `json:"file,omitempty" xorm:"file" db:"file"`
	User      string `json:"user,omitempty" xorm:"user" db:"user"`
	Remote    string `json:"remote,omitempty" xorm:"remote" db:"remote"`
	Uri       string `json:"uri,omitempty" xorm:"uri" db:"uri"`
	UserAgent string `json:"user_agent,omitempty" xorm:"user_agent" db:"user_agent"`
	Method    string `json:"method,omitempty" xorm:"method" db:"method"`
	RequestId string `json:"request_id,omitempty" xorm:"request_id" db:"request_id"`
	TraceId   string `json:"trace_id" xorm:"trace_id" db:"trace_id"`
	Metric    string `json:"metric" xorm:"metric" db:"metric"`     //度量指标
	Tag       string `json:"tag,omitempty" xorm:"tag" db:"tag"`    //度量指标的标记
	From      string `json:"from,omitempty" xorm:"from" db:"from"` //来源,修改时使用
	To        string `json:"to,omitempty" xorm:"to" db:"to"`       //目标
	Error     error  `json:"error,omitempty" xorm:"error" db:"error"`
	TracePid  string `json:"trace_pid,omitempty" xorm:"trace_pid" db:"trace_pid"`
}

func (t *Data) BuildCaller() {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return
	}
	f := runtime.FuncForPC(pc)
	t.FuncName = f.Name()
	t.Line = line
	t.File = file

}
func (t *Data) BuildHTTPRequest(req *http.Request) {
	remoteAddr := req.Header.Get("X-Forwarded-For")
	if len(remoteAddr) == 0 {
		remoteAddr = req.RemoteAddr
	}

	t.User = req.Header.Get("username")
	t.Remote = remoteAddr
	t.Uri = req.RequestURI
	t.UserAgent = req.UserAgent()
	t.Method = req.Method
	t.RequestId = req.Header.Get("request_id")

}
func (t *Data) SetError(err error) *Data {
	t.Error = err
	return t
}
func (t *Data) SetTag(tag string) *Data {
	t.Tag = tag
	return t
}
func (t *Data) SetFrom(from string) *Data {
	t.From = from
	return t
}
func (t *Data) SetTo(to string) *Data {
	t.To = to
	return t
}
func (t *Data) Marshal() []byte {
	b, _ := json.MarshalIndent(t, "", "  ")
	return b
}
