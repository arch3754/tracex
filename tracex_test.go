package tracex

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/arch3754/toolbox/kafka/producer"
	"github.com/arch3754/toolbox/mysql"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"testing"

	"github.com/arch3754/tracex/common"
	"github.com/arch3754/tracex/plugin/kafkawriter"
	"github.com/arch3754/tracex/plugin/xormwriter"
)

func TestXormNew(t *testing.T) {
	c, err := mysql.NewMySQLClient(&mysql.MySQLConf{
		Addr:  "root:1234qwer@tcp(127.0.0.1:3306)/EdgeMgmt?charset=utf8&parseTime=True&loc=Asia%2FShanghai",
		Max:   16,
		Idle:  1,
		Debug: false,
	})
	if err != nil {
		t.Error(err)
		return
	}

	req, _ := http.NewRequest("GET", "http://1.1.1.1/api/location", nil)
	req.Header.Set("username", "root")
	req.Header.Set("request_id", uuid.NewV1().String())
	eng := xormwriter.NewXormWriter(c.Engine, "")
	tx := NewRequestAndCallerTraceX(req, "req", eng.Clone())
	defer tx.Flush()
}

func TestKafka(t *testing.T) {
	c, err := producer.NewProducer(producer.KafkaConf{
		Topic:           "trace",
		BrokersPeers:    "192.168.3.1:9092,192.168.3.2:9092,192.168.3.3:9092",
		SaslUser:        "",
		SaslPasswd:      "",
		Retry:           3,
		DialTimeout:     "30s",
		KeepAlive:       "30s",
		MaxMessageBytes: 209648,
		Version:         sarama.V0_10_2_0,
	})
	if err != nil {
		t.Error(err)
		return
	}

	req, _ := http.NewRequest("GET", "http://1.1.1.1/api/location", nil)
	req.Header.Set("username", "root")
	req.Header.Set("request_id", uuid.NewV1().String())
	eng := kafkawriter.NewKafkaWriter(c)
	tx := NewRequestAndCallerTraceX(req, "req", eng.Clone())
	defer tx.Flush()
}
func TestArr(t *testing.T) {
	var arr []*common.Data
	var a = &common.Data{RequestId: uuid.NewV1().String()}
	arr = append(arr, a)
	a.SetTag("11111").SetTo("222").SetFrom("333")
	a.Method = "GET"
	fmt.Println(arr[0])
}
