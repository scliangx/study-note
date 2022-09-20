package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func NewKafka() *SaramaKafka {
	return &SaramaKafka{
		brokers:           BrokerServer,
		topics:            Topics,
		group:             Group,
		channelBufferSize: 1000,
		ready:             make(chan bool),
		version:           "2.8.0",
		assignor:          Assignor,
	}
}

// Connect 建立连接
func (k *SaramaKafka) Connect() func() {
	fmt.Println("kafka init...")

	version, err := sarama.ParseKafkaVersion(k.version)
	if err != nil {
		_ = fmt.Errorf("error parsing Kafka version: %v", err)
	}

	config := sarama.NewConfig()
	config.Version = version
	// 分区分配策略
	switch Assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", Assignor)
	}
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.ChannelBufferSize = k.channelBufferSize // channel长度

	// 创建client
	newClient, err := sarama.NewClient(BrokerServer, config)
	if err != nil {
		panic(err)
	}
	// 获取所有的topic
	topics, err := newClient.Topics()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("topics: ", topics)

	// 根据client创建consumerGroup
	client, err := sarama.NewConsumerGroupFromClient(k.group, newClient)
	if err != nil {
		log.Fatalf("Error creating consumer group client: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			err := client.Consume(ctx, k.topics, k)
			if err != nil {
				// 当setup失败的时候，error会返回到这里
				_ = fmt.Errorf("error from consumer: %s", err)
				return
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				fmt.Println(ctx.Err())
				return
			}
			k.ready = make(chan bool)
		}
	}()
	<-k.ready
	// 保证在系统退出时，通道里面的消息被消费
	return func() {
		fmt.Println("kafka close")
		cancel()
		wg.Wait()
		if err = client.Close(); err != nil {
			_ = fmt.Errorf("Error closing client: %v \n", err)
		}
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (k *SaramaKafka) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Println("setup")
	//session.ResetOffset(k.topics, 0, 13, "")
	fmt.Println(session.Claims())
	close(k.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (k *SaramaKafka) Cleanup(sarama.ConsumerGroupSession) error {
	fmt.Println("cleanup")
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (k *SaramaKafka) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// 具体消费消息
	for message := range claim.Messages() {
		fmt.Printf("[topic:%s] [partiton:%d] [offset:%d] [value:%s] [time:%v] \n",
			message.Topic, message.Partition, message.Offset, string(message.Value), message.Timestamp)
		// 更新位移
		session.MarkMessage(message, "")
	}
	return nil
}

func main() {
	k := NewKafka()
	c := k.Connect()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		fmt.Println("terminating: via signal")
	}
	c()
}
