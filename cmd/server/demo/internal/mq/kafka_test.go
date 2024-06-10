package mq

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testMq(t *testing.T) {
	const topic = "test_consumer_group_rebalance_test_topic"
	InitSyncProducer()
	InitConsumerGroup()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("init kafka..."),
	}
	_, _, err := SyncProducer.SendMessage(msg)
	assert.NoError(t, err)

	for {
		err := ConsumerGroup.Consume(context.Background(), []string{"my-topic"}, exampleConsumerGroupHandler{})
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}

	_ = ConsumerGroup.Close()

}

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		sess.MarkMessage(msg, "")
	}
	return nil
}
