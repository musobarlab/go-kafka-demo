package sub

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

//SubscriberImpl struct
type SubscriberImpl struct {
	c sarama.Consumer
}

//NewSubscriber constructor of SubscriberImpl
func NewSubscriber(address string) (*SubscriberImpl, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{address}, config)
	if err != nil {
		panic(err)
	}

	return &SubscriberImpl{consumer}, nil
}

//Subscribe function
func (s *SubscriberImpl) Subscribe(topic string) {
	consumer, err := s.c.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Println(err)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	forever := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				fmt.Println("Received messages", string(msg.Key))

				var message Message
				_ = json.Unmarshal(msg.Value, &message)
				fmt.Println("Topic : ", msg.Topic)
				fmt.Println("-------------")
				fmt.Println(message)
			case <-signals:
				fmt.Println("Interrupt is detected")
				forever <- struct{}{}
			}
		}
	}()
	<-forever
}
