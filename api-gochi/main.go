package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	//"github.com/brunodelucasbarbosa/learning-go/api-gochi/repository/userRepository"

	userService "github.com/brunodelucasbarbosa/learning-go/api-gochi/src/services"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/users", userService.FindAllUsers)
	r.Get("/user/{id}", userService.FindUserById)
	r.Get("/user/kafka", SendMessageKafka)

	r.Post("/user", userService.CreateUser)
	r.Patch("/user/{id}", userService.UpdateUser)

	r.Delete("/user/{id}", userService.DeleteUser)

	fmt.Println("Server started on port 3000......")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", r))
}

func SendMessageKafka(w http.ResponseWriter, r *http.Request) {
	//PRODUCER
	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "new-topic",
	}

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte("Hello World"),
		Headers: []protocol.Header{
			{
				Key:   "session",
				Value: []byte("123"),
			},
		},
	})

	if err != nil {
		log.Fatal("Error writing message to kafka")
	}

	//CONSUMER
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "consumer",
		Topic:    "new-topic",
		MinBytes: 0,
		MaxBytes: 10e6, //10MB
	})

	for i := 0; i < 1; i++ {
		message, err := reader.ReadMessage(context.Background())

		for _, val := range message.Headers {
			if val.Key == "session" && string(val.Value) == "123" {
				fmt.Println("session correct")
			}
		}

		if err != nil {
			log.Fatal("Error reading message from kafka", err)
			reader.Close()
		}

		fmt.Print("Message: ", string(message.Value))
	}
	reader.Close()
}
