package kafka

import (
	"github.com/asim/go-micro/plugins/broker/kafka/v3"
	"github.com/asim/go-micro/v3/broker"
	"log"
)

//TODO - UN CONSUMER solo CON UN CASE DEPENDIENDO LA COLA, EL MENSAJE LO TRANSFORMA EN UN OBJETO AL CONSUMIRLO

func ConsumeMessages(topic string) {
	brk := kafka.NewBroker()
	if err := brk.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	_, err := brk.Subscribe(topic, func(event broker.Event) error {
		log.Printf("Received message: %s", string(event.Message().Body))
		return nil
	})

	if err != nil {
		log.Fatalf("Subscribe error: %v", err)
	}

	select {} // block forever
}
