package kafkawriter

import (
	"github.com/arch3754/toolbox/kafka/producer"
	"log"
	"sync"

	"github.com/arch3754/tracex/common"
)

type KafkaWriter struct {
	client *producer.Sender
	datas  []*common.Data
	once   *sync.Once
}

func NewKafkaWriter(sender *producer.Sender) *KafkaWriter {
	return &KafkaWriter{client: sender, once: new(sync.Once)}
}
func (x *KafkaWriter) Set(trace *common.Data) {
	x.datas = append(x.datas, trace)
}
func (x *KafkaWriter) Flush() {
	x.once.Do(func() {
		for _, v := range x.datas {
			if err := x.client.Send(v); err != nil {
				log.Printf("[KafkaWriter] ERROR:flush %v", err)
				return
			}
		}
	})
}
func (x *KafkaWriter) Clone() *KafkaWriter {
	return &KafkaWriter{
		client: x.client,
	}
}
