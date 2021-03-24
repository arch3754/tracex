package xormwriter

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"sync"

	"github.com/arch3754/tracex/common"
)

type XormWriter struct {
	session *xorm.Engine
	datas   []*common.Data
	table   string
	once    *sync.Once
}

func NewXormWriter(sess *xorm.Engine, table string) *XormWriter {
	return &XormWriter{session: sess, table: table, once: new(sync.Once)}
}
func (x *XormWriter) Set(trace *common.Data) {
	if trace == nil {
		panic(fmt.Errorf("trace is nil"))
	}
	x.datas = append(x.datas, trace)
}
func (x *XormWriter) Flush() {
	x.once.Do(func() {
		var err error
		if len(x.table) > 0 {
			_, err = x.session.Table(x.table).Insert(x.datas)
		} else {
			_, err = x.session.Insert(x.datas)
		}
		log.Printf("[XormWriter] ERROR:flush %v", err)
	})
	return
}
func (x *XormWriter) Clone() *XormWriter {
	return &XormWriter{
		session: x.session,
		table:   x.table,
	}
}
