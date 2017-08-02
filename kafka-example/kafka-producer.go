package main

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
)

func createKafkaProducer (kafkaConn string) (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionNone
	var err error
	producer, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	if err != nil {
		return nil, err
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	go func() {
		<-c
		if err := producer.Close(); err != nil {
			log.Fatal("Error closing async producer", err)
		}

		log.Println("Async Producer closed")
		os.Exit(1)
	}()

	go func() {
		for err := range producer.Errors() {
			log.Println("Failed to write message to topic:", err)
		}
	}()

	return producer, nil
}