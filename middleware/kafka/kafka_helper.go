package kafka

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

type KafkaProducer struct {
	address      []string
	syncProducer sarama.SyncProducer
}

//实例化kafka
func NewKafkaProducer(address []string) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return nil, err
	}
	return &KafkaProducer{
		address:      address,
		syncProducer: p,
	}, nil
}

//同步发送
func (self *KafkaProducer) SyncSend(data interface{}, topic string) (int32, int64, error) {
	bs, err := json.Marshal(data)
	if err != nil {
		return -1, -1, err
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(bs),
	}
	part, offset, err := self.syncProducer.SendMessage(msg)
	return part, offset, err
}
