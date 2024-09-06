package kafka

import (
	"github.com/asim/go-micro/plugins/broker/kafka/v3"
	"github.com/asim/go-micro/v3/broker"
	"log"
)

const TopicProduceMessageExample1 = " TopicNameProduce.Example1"
const TopicProduceMessageExample2 = " TopicNameProduce.Example2"

// PublishWhateverExample1 Starts with a capital letter because it's a public function
func PublishWhateverExample1(message []byte) {
	// TODO - seguro aca si me fijo en el de java, este metodo recibia una clase X, que despues la convertia a mensaje de queue
	// TODO - me faltaria cambiar eso
	produceMessage(TopicProduceMessageExample1, message)
}

// PublishWhateverExample2 Starts with a capital letter because it's a public function
func PublishWhateverExample2(message []byte) {
	produceMessage(TopicProduceMessageExample2, message)
}

// produceMessage Starts with a lower case letter because it's a private function
func produceMessage(topic string, message []byte) {
	brk := kafka.NewBroker()
	if err := brk.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	msg := &broker.Message{
		Header: map[string]string{
			"Id": "1",
		},
		Body: message,
	}

	if err := brk.Publish(topic, msg); err != nil {
		log.Fatalf("Publish error: %v", err)
	}
}
