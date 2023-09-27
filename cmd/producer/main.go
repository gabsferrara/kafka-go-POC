package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()
	Publish("messagemxx", "teste", producer, nil, deliveryChan)
	go DeliveryReport(deliveryChan)
	// e := <-deliveryChan

	// msg := e.(*kafka.Message)
	// if msg.TopicPartition.Error != nil {
	// 	fmt.Println("Erro ao enviar")
	// } else {
	// 	fmt.Println("Mesangem enviada: ", msg.TopicPartition)
	// }

	producer.Flush(1000) // p dar tempo de enviar a mensagem

}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka-poc-kafka-1:9092",
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

// infinityloop
func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar")
				// da pra por politicas de retry
			} else {
				fmt.Println("Mesangem enviada: ", ev.TopicPartition)
				//da pra por checkbox
			}
		}
	}
}
