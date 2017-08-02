package main

import (
	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
	"log"
	"os"
	"os/signal"
	"time"
)

const consumerGroup = "group.testing"

type messageHandler func(*sarama.ConsumerMessage) error

func consumeMessages(zookeeperConn string, handler messageHandler) {
	log.Println("Starting Consumer")
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	consumer, err := consumergroup.JoinConsumerGroup(consumerGroup, []string{topicName}, []string{zookeeperConn}, config)
	if err != nil {
		log.Fatal("Failed to join consumer group", consumerGroup, err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	go func() {
		<-c
		if err := consumer.Close(); err != nil {
			log.Println("Error closing the consumer", err)
		}

		log.Println("Consumer closed")
		os.Exit(0)
	}()

	go func() {
		for err := range consumer.Errors() {
			log.Println(err)
		}
	}()

	log.Println("Waiting for messages")
	for message := range consumer.Messages() {
		log.Printf("Topic: %s\t Partition: %v\t Offset: %v\n", message.Topic, message.Partition, message.Offset)

		e := handler(message)
		if e != nil {
			log.Fatal(e)
			consumer.Close()
		} else {
			consumer.CommitUpto(message)
		}
	}
}
