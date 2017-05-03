package queue

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"pgu_status/types"
)

// Listener has msgs
type Listener struct {
	msgs           <-chan amqp.Delivery
	conn           *amqp.Connection
	ch             *amqp.Channel
	errorQueueName string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// NewListener with settings from config.toml
func NewListener(connStr string, queueName string, errorQueueName string) *Listener {
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
	failOnError(err, "Failed to declare a Result queue")

	_, err2 := ch.QueueDeclare(
		errorQueueName, // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err2, "Failed to declare a Error queue")

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

	return &Listener{msgs, conn, ch, errorQueueName}
}

// QueueError Запускает очередь на прослушивание
func (listener Listener) QueueError(d *amqp.Delivery) {
	err := listener.ch.Publish(
		"", // exchange
		listener.errorQueueName, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: d.ContentType,
			Body:        d.Body,
		})
	if err != nil {
		panic(err)
	}
	d.Ack(false)
}

// Start Запускает очередь на прослушивание
func (listener Listener) Start(parser types.IResultParser, finder types.ITaskFinder, sxService types.ISxService) {
	forever := make(chan bool)

	go func() {
		for d := range listener.msgs {
			log.Printf("[INFO] Received a message")
			msg, err := parser.Parse(d.Body)
			if err != nil {
				log.Printf("[ERROR] Parse exeption: '%v'. Body:\n'%v'", err, d.Body)
				listener.QueueError(&d)
				continue
			}
			log.Printf("[INFO] Parse a message (UmmsID: '%v')", msg.UmmsID())

			msg2, err := finder.Find(msg.UmmsID())
			if err != nil {
				log.Printf("[ERROR] Find in SX DB exeption: %v", err)
				listener.QueueError(&d)
				continue
			}
			log.Printf("[INFO] Find SX Task (UmmsID: '%v'. MessageID: '%v')", msg.UmmsID(), msg2.ExtNumber())

			msg3 := types.MakePguStatusMsg(msg, msg2)
			err = sxService.ChangePguCaseStatus(msg3)
			if err != nil {
				log.Printf("[ERROR] Update PGU Case exeption: %v", err)
				listener.QueueError(&d)
				continue
			}
			log.Printf("[INFO] UPDATE CASE on PGU (UmmsID: '%v'. Comment: '%v')", msg.UmmsID(), msg3.Comment())

			d.Ack(false)
			log.Printf("[INFO] Ack message (UmmsID: '%v')", msg.UmmsID())
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
