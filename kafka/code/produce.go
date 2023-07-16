package produce

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaRequest struct {
	Username string
	Age      int
	Email    string
	Des      string
}

func buildRequest(des string) KafkaRequest {
	return KafkaRequest{
		Username: "golang001",
		Age:      10,
		Email:    "golang001@golang.com",
		Des:      des,
	}
}

func InitializationProducer() {
	config := kafka.ConfigMap{
		"bootstrap.servers":   "127.0.0.1:9092,127.0.0.1:9093,127.0.0.1:9094",
		"api.version.request": "true",
		"message.max.bytes":   1000000,
		"linger.ms":           10,
		"retries":             30,
		"retry.backoff.ms":    1000,
		"acks":                "1",
		"security.protocol":   "plaintext",
	}

	// 创建生产者
	producer, err := kafka.NewProducer(&config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v Message: %v\n", ev.TopicPartition, string(ev.Value))
				}
			}
		}
	}()
	topic := "first"

	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		msg := buildRequest(word)
		byteData, _ := json.Marshal(msg)
		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: 0},
			Value:          byteData,
		}, nil)
		if err != nil {
			fmt.Println("error", err)
		}
	}

	// Wait for message deliveries before shutting down
	producer.Flush(60 * 1000)
}
