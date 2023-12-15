package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

func main() {
	e := echo.New()
	kConf, err := kafkaWriter("message-to-say")
	if err != nil {
		panic(err)
	}
	e.GET("/says/:msg", func(c echo.Context) error {
		msg := c.Param("msg")
		ctx := c.Request().Context()
		publish(ctx, kConf, nil, []byte(msg))
		fmt.Printf("published: %s\n", msg)
		return c.String(http.StatusOK, msg)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func publish(ctx context.Context, w *kafka.Writer, key, value []byte) (err error) {
	message := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}
	return w.WriteMessages(ctx, message)
}

func kafkaWriter(topic string) (w *kafka.Writer, err error) {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: "kafka",
	}

	config := kafka.WriterConfig{
		Brokers:          []string{"127.0.0.1:9093"},
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	w = kafka.NewWriter(config)
	return w, nil
}
