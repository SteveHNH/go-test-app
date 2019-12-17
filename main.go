package main

import (
	"context"
	"fmt"

	"github.com/SteveHNH/go-test-app/config"
	l "github.com/SteveHNH/go-test-app/logger"
	"github.com/segmentio/kafka-go"
)

func main() {
	cfg := config.Get()
	l.InitLogger()

	consumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers: cfg.KafkaBrokers,
		GroupID: cfg.KafkaGroupID,
		Topic:   cfg.ConsumeTopic,
	})

	producer := kafka.NewWriter(Kafka.WriterConfig{
		Brokers:  cfg.KafkaBrokers,
		Topic:    cfg.ProduceTopic,
		Balancer: &kafka.LeastBytes{},
	})

	for {
		m, err := consumer.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("Message contents: %s", string(m.Value))
	}

}
