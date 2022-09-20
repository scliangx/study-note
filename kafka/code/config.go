package main

type SaramaKafka struct {
	brokers           []string
	topics            []string
	startOffset       int64
	version           string
	ready             chan bool
	group             string
	channelBufferSize int
	assignor          string
}

var (
	Topics = []string{"myTopic"}
	// BrokerServer 集群地址
	BrokerServer = []string{"8.141.175.100:9092", "8.141.175.100:9093", "8.141.175.100:9094"}
	Assignor     = "range"
	Group        = "myGroup"
)
