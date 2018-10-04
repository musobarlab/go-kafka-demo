package pub

import (
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

//PublisherImpl struct
type PublisherImpl struct {
	producer sarama.SyncProducer
}

//NewPublisher constructor of PublisherImpl
func NewPublisher(address string) (*PublisherImpl, error) {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// producer config
	config := sarama.NewConfig()
	config.ClientID = "WURY"
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	//config.Metadata.Retry.Backoff = 2 * time.Second

	// async producer
	//prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{address}, config)

	if err != nil {
		return nil, err
	}

	return &PublisherImpl{prd}, nil
}

//Publish function
func (publiher *PublisherImpl) Publish(topic string, message []byte) error {
	// publish sync
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}
	p, o, err := publiher.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Println("Partition ", p)
	fmt.Println("Offset ", o)

	// publish async
	//producer.Input() <- &sarama.ProducerMessage{
	return nil
}
