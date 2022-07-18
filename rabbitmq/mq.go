package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	ch       *amqp.Channel
	Name     string
	exchange string
}

func New(s string) *RabbitMQ {
	conn, e := amqp.Dial(s)
	if e != nil {
		panic(e)
	}
	ch, e := conn.Channel()
	if e != nil {
		panic(e)
	}
	q, e := ch.QueueDeclare(
		"",    // name
		false, // durable
		true,  // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if e != nil {
		panic(e)
	}
	return &RabbitMQ{
		ch:   ch,
		Name: q.Name,
	}
}

func (mq *RabbitMQ) Bind(exchange string) {
	e := mq.ch.QueueBind(
		mq.Name,  // queue name
		"",       // routing key
		exchange, // exchange
		false,
		nil,
	)
	if e != nil {
		panic(e)
	}
	mq.exchange = exchange
}

func (mq *RabbitMQ) Send(key string, body interface{}) {
	bytes, e := json.Marshal(body)
	if e != nil {
		panic(e)
	}
	e = mq.ch.Publish(
		"",  // exchange
		key, // routing key
		false,
		false,
		amqp.Publishing{
			Body:    bytes,
			ReplyTo: mq.Name,
		},
	)
	if e != nil {
		panic(e)
	}
}

func (mq *RabbitMQ) Publish(exchange string, body interface{}) {
	bytes, e := json.Marshal(body)
	if e != nil {
		panic(e)
	}
	e = mq.ch.Publish(
		exchange, // exchange
		"",       // routing key
		false,
		false,
		amqp.Publishing{
			Body:    bytes,
			ReplyTo: mq.Name,
		},
	)
	if e != nil {
		panic(e)
	}
}

func (mq *RabbitMQ) Consume() <-chan amqp.Delivery {
	c, e := mq.ch.Consume(
		mq.Name, // queue
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	if e != nil {
		panic(e)
	}
	return c
}

func (mq *RabbitMQ) Close() {
	if e := mq.ch.Close(); e != nil {
		panic(e)
	}
}
