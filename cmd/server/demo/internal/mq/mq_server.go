package mq

import (
	"github.com/IBM/sarama"
	"github.com/aide-family/moon/cmd/server/demo/internal/democonf"
	"github.com/aide-family/moon/pkg/mq/kafka/consume"
	"github.com/aide-family/moon/pkg/mq/kafka/producer"
	"github.com/go-kratos/kratos/v2/log"
)

type KafkaMQ struct {
	log           *log.Helper
	consumerGroup sarama.ConsumerGroup
	producer      sarama.SyncProducer
}

func InitKafkaMQServer(mqConf democonf.KafkaMQServerConfig) *KafkaMQ {

	kafkaEndpoints := mqConf.GetEndpoints()
	kafkaGroupId := mqConf.GetGroupId()
	username := mqConf.GetUsername()
	password := mqConf.GetPassword()

	conf := &sarama.Config{}

	if len(username) > 0 || len(password) > 0 {
		conf.Net.SASL.Enable = true
		conf.Net.SASL.User = username
		conf.Net.SASL.Password = password
	}

	initProducer := InitProducer(kafkaEndpoints, conf)
	consumerGroup := InitConsumerGroup(kafkaEndpoints, kafkaGroupId, conf)

	return &KafkaMQ{
		//log:           log.NewHelper(log.With(logger, "module", "demo.kafka_mq")),
		consumerGroup: consumerGroup,
		producer:      initProducer,
	}
}

func InitProducer(brokers []string, conf *sarama.Config) sarama.SyncProducer {
	conf.Net.MaxOpenRequests = 5
	conf.Producer.MaxMessageBytes = 1000000
	conf.Producer.RequiredAcks = sarama.RequiredAcks(1)
	conf.Producer.Retry.Max = 1
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	return producer.InitSynProducer(brokers, conf)
}

func InitConsumerGroup(brokers []string, kafkaGroupID string, conf *sarama.Config) sarama.ConsumerGroup {
	return consume.ConsumerGroup(brokers, kafkaGroupID, conf)
}
