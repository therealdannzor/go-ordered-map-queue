package queue

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/therealdannzor/go-ordered-map-queue/io"
	"github.com/therealdannzor/go-ordered-map-queue/model/command"
	"github.com/therealdannzor/go-ordered-map-queue/model/ordmap"
)

type Server struct {
	store ordmap.OrdMap
}

func NewServer() Server {
	return Server{
		store: ordmap.New(),
	}
}

func (s *Server) Receive() error {
	conn, err := amqp.Dial(servUrl)
	if err != nil {
		return fmt.Errorf("consumer failed to dial server url: %w", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("consumer failed to open channel: %w", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"ord-map-queue",
		false,
		false,
		false, // not exclusive (multiple clients)
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("consumer failed to declare queue: %w", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false, // not exclusive (multiple clients)
		false,
		false,
		nil,
	)

	var persistentCh chan struct{}

	go func() {
		for m := range msgs {
			cmd, err := unpack(m.Body)
			if err != nil {
				fmt.Println("failed to read body")
			}

			switch cmd.Operation {
			case command.Add:
				s.store.Update(cmd.Params[0], cmd.Params[1])
			case command.Delete:
				s.store.Delete(cmd.Params[0])
			case command.Lookup:
				res := s.store.Get(cmd.Params[0])
				if err := io.WriteFile([]byte(res)); err != nil {
					fmt.Println("Failed to write to file: ", err)
				}
			case command.GetAll:
				res := s.store.GetAll()
				if err := io.WriteFile([]byte(fmt.Sprint(res))); err != nil {
					fmt.Println("Failed to write to file: ", err)
				}
			default:
				fmt.Println("No-op: this should not happen")
			}

			m.Ack(true) // ack every delivery
		}
	}()

	log.Printf("Waiting for messages..")
	<-persistentCh

	return nil
}

func unpack(payload []byte) (*command.Command, error) {
	var result command.Command
	if err := json.Unmarshal(payload, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
