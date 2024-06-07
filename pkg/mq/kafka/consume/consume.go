package consume

import (
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
)

type Consumer struct {
	ready chan bool
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/IBM/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Info("message channel was closed")
				return nil
			}
			log.Infof("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}

func ConsumerGroup(brokers []string, groupId string, conf *sarama.Config) sarama.ConsumerGroup {
	log.Infof("kafka comsumerGroup brokers: %v  groupId:%v conf: %v", brokers, groupId, conf)
	cGroup, err := sarama.NewConsumerGroup(brokers, groupId, conf)
	if err != nil {
		panic(err)
	}
	return cGroup
}
