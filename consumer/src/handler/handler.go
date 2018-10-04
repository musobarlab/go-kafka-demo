package handler

import (
	"github.com/wuriyanto48/go-kafka-demo/consumer/src/sub"
)

//WorkerHandler struct
type WorkerHandler struct {
	topic      string
	subscriber sub.Subscriber
}

//NewWorkerHandler WorkerHandler's constructor
func NewWorkerHandler(topic string, subscriber sub.Subscriber) *WorkerHandler {
	return &WorkerHandler{topic: topic, subscriber: subscriber}
}

//Pool function
func (h *WorkerHandler) Pool() {
	h.subscriber.Subscribe(h.topic)
}
