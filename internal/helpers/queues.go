package helpers

import (
	"fmt"
	"encoding/json"
	"bytes"
	"github.com/streadway/amqp"
)

type Message struct{
	Data	interface{}	
	Event	QueueEvent
}

type QueueEvent struct{
	Name	string `json:"name"`
	Options	*map[string]interface{} `json:"options"`
}

func MQConn(mqUser,mqPassword,mqHost,mqPort string) (*amqp.Connection, error){
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/",mqUser,mqPassword,mqHost,mqPort))
	
	return conn, err
}

func MQChannel(mqUser, mqPassword, mqHost, mqPort string) (*amqp.Connection, *amqp.Channel, error){
	conn, err := MQConn(mqUser,mqPassword,mqHost,mqPort)	

	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()

	return conn, ch, err
}

func MQQueue(qName string, ch *amqp.Channel) (amqp.Queue, error){
	q, err := ch.QueueDeclare(
		qName,
		false,
		false,
		false,
		false,
		nil,
	)

	return q,err
}

func MQPublish(ch *amqp.Channel, q amqp.Queue, body string) error {
	err := ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		},
	)	

	return err
}

func MQPublishMessage(ch *amqp.Channel, q amqp.Queue, message Message) error {
	body, err := message.Serialize()

	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: body,
		},
	)	

	return err
}

func MQConsumer(ch *amqp.Channel, q amqp.Queue)(<-chan amqp.Delivery,error){
	consumer, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	return consumer, err
}

func (m *Message)Serialize() ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(m)
	return b.Bytes(), err
}

func (m *Message)Deserialize(b []byte) error {
	buffer := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buffer)
	err := decoder.Decode(m)
	return err
}


