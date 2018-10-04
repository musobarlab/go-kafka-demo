package main

import (
	"fmt"
	"os"

	configEnv "github.com/joho/godotenv"
	"github.com/wuriyanto48/go-kafka-demo/consumer/src/handler"
	"github.com/wuriyanto48/go-kafka-demo/consumer/src/sub"
)

func main() {
	fmt.Println("consumer")

	err := configEnv.Load(".env")

	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	zookeeperHost, ok := os.LookupEnv("ZOOKEEPER_HOST")

	if !ok {
		fmt.Println("cannot load ZOOKEEPER_HOST from environment")
		os.Exit(2)
	}

	topic, ok := os.LookupEnv("KAFKA_TOPIC")

	if !ok {
		fmt.Println("cannot load KAFKA_TOPIC from environment")
		os.Exit(2)
	}

	subscriber, err := sub.NewSubscriber(zookeeperHost)

	if err != nil {
		fmt.Println("error create subscriber")
		os.Exit(2)
	}

	workerHandler := handler.NewWorkerHandler(topic, subscriber)

	workerHandler.Pool()
}
