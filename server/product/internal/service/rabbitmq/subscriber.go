package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"github.com/NetSinx/yconnect-shop/server/product/db"
	"github.com/NetSinx/yconnect-shop/server/product/errs"
	"github.com/NetSinx/yconnect-shop/server/product/model"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm/clause"
)

func ConsumeCategoryEvents() {
	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%s",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)
	conn, err := amqp.Dial(amqpURL)
	errs.LogOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	errs.LogOnError(err, "Failed to open a channel")
	defer ch.Close()

	exchange := "category_events"
	err = ch.ExchangeDeclare(
		exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	errs.LogOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"product-service-category-mirror", true, false, false, false, nil,
	)
	errs.LogOnError(err, "Failed to declare a queue")

	bindings := []string{"category.created", "category.updated", "category.deleted"}
	for _, rk := range bindings {
		err := ch.QueueBind(
			q.Name,
			rk,
			exchange,
			false,
			nil,
		)
		errs.LogOnError(err, "Failed to binding a queue")
	}

	msgs, err := ch.Consume(
		q.Name, "", true, false, false, false, nil,
	)
	errs.LogOnError(err, "Failed to consume messages")
	log.Println("Consuming category events...")

	db := db.ConnectDB()

	for d := range msgs {
		rk := d.RoutingKey
		var categoryMirror model.CategoryMirror
		if err := json.Unmarshal(d.Body, &categoryMirror); err != nil {
			log.Println("Invalid message:", err)
			continue
		}

		switch rk {
		case "category.created":
			db.Clauses(clause.OnConflict{DoNothing: true}).Create(&categoryMirror)
		case "category.updated":
			db.Where("id = ?", categoryMirror.Id).Updates(&categoryMirror)
		case "category.deleted":
			db.Delete(&categoryMirror, "slug = ?", categoryMirror.Slug)
		default:
		}
	}
}