package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	// "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// InitProducer()
	AsyncProducer()
}

// InitProducer 同步生产
func InitProducer() {
	config := sarama.NewConfig()
	// request.timeout.ms
	config.Producer.Timeout = time.Second * 5
	// message.max.bytes
	config.Producer.MaxMessageBytes = 1024 * 1024
	// request.required.acks
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Version = sarama.V2_8_1_0
	config.Producer.Partitioner = sarama.NewHashPartitioner

	if err := config.Validate(); err != nil {
		panic(fmt.Errorf("invalid configuration, error: %v", err))
	}

	producer, err := sarama.NewSyncProducer(BrokerServer, config)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	for {
		msg := &sarama.ProducerMessage{
			Topic: Topics[0],
			Value: sarama.StringEncoder("testing first topic..."),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("send message failed: %s\n", err)
		} else {
			log.Printf("message sent to partition %d at offset %d\n", partition, offset)
		}
		time.Sleep(1 * time.Second)
	}
}

// AsyncProducer 异步生产
func AsyncProducer() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true //必须有这个选项
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewAsyncProducer(BrokerServer, config)
	defer p.Close()
	if err != nil {
		return
	}
	//保证通道不会被堵塞
	go func(p sarama.AsyncProducer) {
		errors := p.Errors()
		success := p.Successes()
		for {
			select {
			case err := <-errors:
				if err != nil {
					fmt.Errorf("[ERROR]: %s", err)
				}
			case <-success:
			}
		}
	}(p)

	for {
		v := "async testing message...: " + strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10000))
		msg := &sarama.ProducerMessage{
			Topic: Topics[0],
			Value: sarama.ByteEncoder(v),
		}
		p.Input() <- msg
		time.Sleep(time.Second * 1)
		fmt.Printf("send %v to kafka success\n", v)
	}
}
