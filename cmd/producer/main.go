package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()
	Publish("transferiu2", "teste", producer, []byte("transferencia"), deliveryChan)
	go DeliveryReport(deliveryChan)

	producer.Flush(2000) // p dar tempo de enviar a mensagem

}

func NewKafkaProducer() *kafka.Producer {
	// configuration lib --> https://github.com/confluentinc/librdkafka/blob/master/CONFIGURATION.md

	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "kafka-poc-kafka-1:9092",
		"delivery.timeout.ms": "0",
		"acks":                "all",
		"enable.idempotence":  "true",
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
