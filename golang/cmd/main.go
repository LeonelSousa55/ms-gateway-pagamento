package main

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/codeedu/imersao5-gateway/adapter/broker/kafka"
	"github.com/codeedu/imersao5-gateway/adapter/factory"
	"github.com/codeedu/imersao5-gateway/adapter/presenter/transaction"
	"github.com/codeedu/imersao5-gateway/usecase/process_transaction"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
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
		"bootstrap.servers": "host.docker.internal:9094",
	}
	kafkaPresenter := transaction.NewTransactionKafkaPresenter()

	//Config. producer
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)
	var msgChan = make(chan *ckafka.Message)

	//Config. configMapConsumer
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
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
