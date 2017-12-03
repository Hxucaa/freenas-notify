package mq

import (
	"github.com/streadway/amqp"
	"github.com/hxucaa/freenas-notify/util"
)

func send(url string, key string, msg string) {


	conn, err := amqp.Dial(url)
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	util.FailOnError(err, "Failed to declare an exchange")

	err = ch.Publish(
		exchangeName,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(msg),
			UserId: userId,
		},
	)
	util.FailOnError(err, "Failed to publish a message")
}

func SendPostInit(url string) {
	send(url, routingKey, "postinit")
}

func SendPreInit(url string) {
	send(url, routingKey, "preinit")
}

func SendShutdown(url string) {
	send(url, routingKey, "shutdown")
}