package kafkaconnect

import (
	"log"

	"github.com/dilshodforever/5-oyimtixon/kafka"
)

func ConnectKAfka() kafka.KafkaProducer{
	kaf, err := kafka.NewKafkaProducer([]string{"localhost:9092"})
	if err != nil {
		log.Fatal("Error while connection kafka: ", err.Error())
	}
	return kaf

}
