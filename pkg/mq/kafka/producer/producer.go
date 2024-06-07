package producer

import (
	"fmt"
	"github.com/IBM/sarama"
	"sync"
)

type topicPartition struct {
	topic     string
	partition int32
}

type producerProvider struct {
	producersLock sync.Mutex
	producers     map[topicPartition][]sarama.AsyncProducer

	producerProvider func(topic string, partition int32) sarama.AsyncProducer
}

func InitSynProducer(brokers []string, conf *sarama.Config) sarama.SyncProducer {
	producer, err := sarama.NewSyncProducer(brokers, conf)
	if err != nil {
		panic(err.Error())
	}
	return producer
}

func newProducerProvider(brokers []string, producerConfigurationProvider func() *sarama.Config) *producerProvider {
	provider := &producerProvider{
		producers: make(map[topicPartition][]sarama.AsyncProducer),
	}
	provider.producerProvider = func(topic string, partition int32) sarama.AsyncProducer {
		config := producerConfigurationProvider()
		if config.Producer.Transaction.ID != "" {
			config.Producer.Transaction.ID = config.Producer.Transaction.ID + "-" + topic + "-" + fmt.Sprint(partition)
		}
		producer, err := sarama.NewAsyncProducer(brokers, config)
		if err != nil {
			return nil
		}
		return producer
	}
	return provider
}

func (p *producerProvider) borrow(topic string, partition int32) (producer sarama.AsyncProducer) {
	p.producersLock.Lock()
	defer p.producersLock.Unlock()

	tp := topicPartition{topic: topic, partition: partition}

	if producers, ok := p.producers[tp]; !ok || len(producers) == 0 {
		for {
			producer = p.producerProvider(topic, partition)
			if producer != nil {
				return
			}
		}
	}

	index := len(p.producers[tp]) - 1
	producer = p.producers[tp][index]
	p.producers[tp] = p.producers[tp][:index]
	return
}

func (p *producerProvider) release(topic string, partition int32, producer sarama.AsyncProducer) {
	p.producersLock.Lock()
	defer p.producersLock.Unlock()

	if producer.TxnStatus()&sarama.ProducerTxnFlagInError != 0 {
		// Try to close it
		_ = producer.Close()
		return
	}
	tp := topicPartition{topic: topic, partition: partition}
	p.producers[tp] = append(p.producers[tp], producer)
}
