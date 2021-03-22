# tracex

##xorm
```
	c, err := mysql.NewMySQLClient(&mysql.MySQLConf{
		Addr:  "root:1234qwer@tcp(127.0.0.1:3306)/trace?charset=utf8&parseTime=True&loc=Asia%2FShanghai",
		Max:   16,
		Idle:  1,
		Debug: false,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
    eng := xormwriter.NewXormWriter(c.Engine, "table")
  
	tx := tracex.NewRequestAndCallerTraceX(req, "req", eng.Clone())
	defer tx.Flush()
```
##kafka
```
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
		fmt.Println(err)
		return
	}
	cli := kafkawriter.NewKafkaWriter(c)
  
	tx := tracex.NewRequestAndCallerTraceX(req, "req", cli.Clone())
	defer tx.Flush()
```