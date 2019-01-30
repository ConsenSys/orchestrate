package main

import (
	flags "github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

// LoggerConfig logger configuration
type LoggerConfig struct {
	Level  string `long:"log-level" env:"LOG_LEVEL" default:"debug" description:"Log level, one of panic, fatal, error, warn, info, debug, trace."`
	Format string `long:"log-format" env:"LOG_FORMAT" default:"text" description:"Log formatter, one of text, json."`
}

// WorkerConfig application configuration
type WorkerConfig struct {
	Slots uint `short:"w" long:"worker-slots" env:"WORKER_SLOTS" default:"100" description:"Number of messages that can be treat in parallel."`
}

// KafkaConfig is the configuration of application dealing with Kafka
type KafkaConfig struct {
	Address       []string `short:"k" long:"kafka-address" env:"KAFKA_ADDRESS" default:"localhost:9092" description:"Address of Kafka server to connect to"`
	InTopic       string   `short:"i" long:"kafka-in-topic" env:"KAFKA_TOPIC_TX_SENDER" default:"topic-tx-sender" description:"Kafka topic to consume message from"`
	ConsumerGroup string   `short:"g" long:"kafka-consumer-group" env:"KAFKA_GROUP_TX_SENDER" default:"group-tx-sender" description:"Kafka consumer group"`
}

// EthConfig is the configuration of application dealing with Ethereum
type EthConfig struct {
	URLs []string `short:"e" long:"eth-client" env:"ETH_CLIENT_URL" default:"https://ropsten.infura.io/v3/81e039ce6c8a465180822b525e3644d7"`
}

// Config worker configuration
type Config struct {
	Log    LoggerConfig
	Worker WorkerConfig
	Kafka  KafkaConfig
	Eth    EthConfig
}

// LoadConfig load configuration
func LoadConfig(opts interface{}) {
	_, err := flags.Parse(opts)
	if err != nil {
		panic(err)
	}
}

// ConfigureLogger configure logger
func ConfigureLogger(opts LoggerConfig) {
	switch opts.Format {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{})
	}
	if logLevel, err := log.ParseLevel(opts.Level); err != nil {
		log.Fatalf("Invalid log level, %v", err)
	} else {
		log.New()
		log.SetLevel(logLevel)
	}
}
