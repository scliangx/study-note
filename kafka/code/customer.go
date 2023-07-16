package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func InitializationConsumer() {
	config := &kafka.ConfigMap{
		"bootstrap.servers":         "127.0.0.1:9092,127.0.0.1:9093,127.0.0.1:9094",
		"group.id":                  "first-group-id1",
		"api.version.request":       "true",
		"auto.offset.reset":         "earliest",
		"heartbeat.interval.ms":     3000,
		"session.timeout.ms":        30000,
		"max.poll.interval.ms":      120000,
		"fetch.max.bytes":           1024000,
		"max.partition.fetch.bytes": 256000,
		"enable.auto.commit":        "false",
	}
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	/*
		topic := "first"
		partition := 0
		offset := kafka.OffsetBeginning
		// 手动分配分区和偏移量
		err = consumer.Assign([]kafka.TopicPartition{
			{
				Topic:     &topic,
				Partition: int32(partition),
				Offset:    offset,
				// 或者
				//Offset:    kafka.OffsetInvalid,
				//Timestamp: timestamp,
			},
		})
	*/

	topics := []string{"first"}
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		panic(err)
	}
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			consumer.CommitOffsets([]kafka.TopicPartition{msg.TopicPartition})
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}

	}
}

