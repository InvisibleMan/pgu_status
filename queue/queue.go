package queue

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"pgu_status/types"
)

// Listener has msgs
type Listener struct {
	msgs <-chan amqp.Delivery
	conn *amqp.Connection
	ch   *amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// NewListener with settings from config.toml
func NewListener(connStr string, queueName string) *Listener {
	conn, err := amqp.Dial(connStr)
	failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	// defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	// log.Printf("[INFO] Starting listen queue '%s'", queueName)

	return &Listener{msgs, conn, ch}
}

// Start Запускает очередь на прослушивание
func (listener Listener) Start(parser types.IResultParser, finder types.ITaskFinder, sxService types.ISxService) {
	forever := make(chan bool)

	go func() {
		for d := range listener.msgs {
			log.Printf("Received a message: %s", d.Body)
			msg, err := parser.Parse(d.Body)
			if err != nil {
				log.Printf("[ERROR] Parse exeption: %v", err)
				continue
			}

			msg2, err := finder.Find(msg.UmmsID())
			if err != nil {
				log.Printf("[ERROR] Find in SX DB exeption: %v", err)
				continue
			}

			log.Printf("[INFO] TRY UPDATE CASE on PGU %v", msg2)
		}
	}()

	<-forever
}

// Close test DB connection
func (listener Listener) Close() {
	if listener.ch != nil {
		listener.ch.Close()
	}
	if listener.conn != nil {
		listener.conn.Close()
	}
}
