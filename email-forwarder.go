package main

import (
	"fmt"
	"os"

	"github.com/catuss-a/imap"
	"github.com/keighl/mandrill"
	"github.com/streadway/amqp"
)

func fatalOnError(err error, msg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

type WorkerConfig struct {
	Environment string

	Amqp struct {
		Hostname     string
		MessageQueue string
	}
	Mandrill struct {
		ClientKey string
		From      string
	}
}

var (
	cfg WorkerConfig
)

const (
	configPath = "./config.yaml"
)

func init() {
	loadConfig(configPath, &cfg)
	fmt.Println("Loading configuration file", configPath)
}

func main() {
	conn, err := amqp.Dial(cfg.Amqp.Hostname)
	fatalOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	fatalOnError(err, "Failed to open a channel")
	defer ch.Close()

	fmt.Println("Queue Name:", cfg.Amqp.MessageQueue)
	q, err := ch.QueueDeclare(
		cfg.Amqp.MessageQueue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	fatalOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	fatalOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	fatalOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		mclient := mandrill.ClientWithKey(cfg.Mandrill.ClientKey)

		fmt.Println("Waiting for new incoming messages to consume...")
		for rawMsg := range msgs {
			fmt.Println("Consuming incoming message...")
			decodedMsg := imapClient.NewMessageFromBytes(rawMsg.Body)
			forwardMessage(mclient, decodedMsg)
			rawMsg.Ack(false)
			fmt.Println("Waiting for new incoming messages to consume...")
		}
	}()
	<-forever
}

func forwardMessage(mclient *mandrill.Client, message *imapClient.GoImapMessage) {
	msg := &mandrill.Message{}

	if cfg.Environment == "development" {
		msg.AddRecipient("axel.catusse@gmail.com", "Axel Catusse", "to")
	} else {
		msg.AddRecipient(message.To, "", "to")
	}

	msg.FromEmail = cfg.Mandrill.From
	msg.Subject = message.Subject
	msg.Text = message.Body

	_, err := mclient.MessagesSend(msg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error when forwarding message", err)
	} else {
		fmt.Println("Message succesfully forwarded")
	}
}
