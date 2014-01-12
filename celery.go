// celery_go project main.go
package main

import (	
	//"fmt"
	"github.com/streadway/amqp",	
	"log"	
)

func main() {
	log.Printf("Starting")
	conn, _ := amqp.Dial("amqp://karma:123456@localhost:5672/")

	channel, _ := conn.Channel()

	deliveries, _ := channel.Consume(
		"celery", //queue
		"asdf_go",
		false,
		false,
		false,
		false,
		nil)

	for d := range deliveries {
		log.Printf("got %s", d.Body)
		d.Ack(false)
	}

	log.Printf("Success")
}
