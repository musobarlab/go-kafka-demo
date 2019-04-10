package pub

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	mp "github.com/wuriyanto48/go-kafka-demo/producer/src/messageproto"
)

//Content struct
type Content struct {
	Header string `json:"header"`
	Body   string `json:"body"`
}

//Message struct
type Message struct {
	From    string  `json:"from"`
	Content Content `json:"content"`
}

//JSON function
func (m *Message) JSON() ([]byte, error) {
	return json.Marshal(m)
}

// ToProto function
func (m *Message) ToProto() ([]byte, error) {
	messageProto := &mp.Message{
		From: m.From,
		Content: &mp.Content{
			Body:   m.Content.Body,
			Header: m.Content.Header,
		},
	}

	return proto.Marshal(messageProto)
}
