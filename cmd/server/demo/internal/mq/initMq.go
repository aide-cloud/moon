package mq

import (
	"flag"
	"github.com/IBM/sarama"
	"github.com/aide-family/moon/pkg/mq/kafka/consume"
	"github.com/aide-family/moon/pkg/mq/kafka/producer"
	"github.com/redis/go-redis/v9"
	"strings"
)

var (
	brokers       = flag.String("brokers", "localhost:9092", "The Kafka brokers to connect to, as a comma separated list")
	RedisClient   *redis.Client
	SyncProducer  sarama.SyncProducer
	ConsumerGroup sarama.ConsumerGroup
)

func InitSyncProducer() {
	splitBrokers := strings.Split(*brokers, ",")
	conf := sarama.NewConfig()
	conf.Net.MaxOpenRequests = 5
	conf.Producer.MaxMessageBytes = 1000000
	conf.Producer.RequiredAcks = sarama.RequiredAcks(1)
	conf.Producer.Retry.Max = 1
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	producer := producer.InitSynProducer(splitBrokers, conf)
	SyncProducer = producer
}

func InitConsumerGroup() {
	splitBrokers := strings.Split(*brokers, ",")
	conf := sarama.NewConfig()
	cGroup := consume.ConsumerGroup(splitBrokers, "aaaa", conf)
	ConsumerGroup = cGroup
}
