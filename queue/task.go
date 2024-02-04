package queue

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/therealdannzor/go-ordered-map-queue/model/command"
)

const servUrl string = "amqp://guest:guest@localhost:5672/"

func Send() error {
	conn, err := amqp.Dial(servUrl)
	if err != nil {
		return fmt.Errorf("producer failed to dial server url: %w", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("producer failed to open channel: %w", err)
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
		return fmt.Errorf("producer failed to declare queue: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		token := scanner.Text()
		line := strings.Split(token, ",")

		cmd, err := parseLine(line)
		if err != nil {
			fmt.Println("invalid instruction with line: ", line, "error: ", err)
		}

		b, err := json.Marshal(cmd)
		if err != nil {
			fmt.Println("failed to marshal cmd of line: ", line, "error: ", err)
		}

		if err = ch.PublishWithContext(
			ctx,
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        b,
			},
		); err != nil {
			fmt.Println("failed to publish: %w", err)
		}
	}

	return nil
}

func parseLine(line []string) (*command.Command, error) {
	op, err := command.ParseOperation(line[0])
	if err != nil {
		return nil, err
	}

	cmd, err := command.ParseCommand(op, line[1:])
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
