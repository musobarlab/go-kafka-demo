package sub

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	mp "github.com/wuriyanto48/go-kafka-demo/consumer/src/messageproto"
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

// UnmarshalProto function
func UnmarshalProto(data []byte) (*Message, error) {
	var m mp.Message
	err := proto.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	var message Message
	message.From = m.From
	message.Content.Header = m.Content.Header
	message.Content.Body = m.Content.Body
	return &message, nil
}
