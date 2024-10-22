package amqp

import (
	"authorization_service/internal/core/configuration"
	"authorization_service/internal/core/models"
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type publisher struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func New(cfg *configuration.AMQPConfig) (*publisher, error) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:5672/", cfg.User, cfg.Pass, cfg.Host)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return &publisher{}, err
	}

	ch, err := conn.Channel()
	return &publisher{ch: ch}, err
}

func (p *publisher) Setup() error {
	return p.ch.ExchangeDeclare("user_update", "fanout", true, false, false, false, nil)
}

func (p *publisher) SendUserUpdateEvent(data models.UpdateUserEvent) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return p.ch.Publish("user_update", "", false, false, amqp.Publishing{
		ContentType: "text/json",
		Body:        []byte(body),
	})
}

func (p *publisher) Close() {
	p.conn.Close()
	p.ch.Close()
}
