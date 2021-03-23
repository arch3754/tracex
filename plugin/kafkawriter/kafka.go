package kafkawriter

import (
	"github.com/arch3754/toolbox/kafka/producer"
	"github.com/arch3754/tracex/common"
)

type KafkaWriter struct {
	client *producer.Sender
	datas  []*common.Data
}

func NewKafkaWriter(sender *producer.Sender) *KafkaWriter {
	return &KafkaWriter{client: sender}
}
func (x *KafkaWriter) Set(trace *common.Data) {
	x.datas = append(x.datas, trace)
}
func (x *KafkaWriter) Flush() error {
	for _, v := range x.datas {
		if err := x.client.Send(v); err != nil {
			return err
		}
	}
	return nil
}
func (x *KafkaWriter) Clone() *KafkaWriter {
	return &KafkaWriter{
		client: x.client,
	}
}
