package main

import (
	"br/com/leonel/adapter/broker/kafka"
	"br/com/leonel/adapter/factory"
	"br/com/leonel/adapter/presenter/transaction"
	"br/com/leonel/usecase/process_transaction"
	"database/sql"
	"encoding/json"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {

	//Config. bd
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	//Config. repository
	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()

	//Config. configMapProducer
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}
	kafkaPresenter := transaction.NewTransactionKafkaPresenter()

	//Config. producer
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)
	var msgChan = make(chan *ckafka.Message)

	//Config. configMapConsumer
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp",
		"group.id":          "goapp",
	}

	//Config. topic
	topics := []string{"transactions"}

	//Config. consumer
	consumer := kafka.NewConsumer(configMapConsumer, topics)
	go consumer.Consume(msgChan)

	//Config. userCase
	usercase := process_transaction.NewProcessTransaction(repository, producer, "transactions_result")

	for msg := range msgChan {
		var input process_transaction.TransactionDtoInput
		json.Unmarshal(msg.Value, &input)
		usercase.Execute(input)
	}
}
