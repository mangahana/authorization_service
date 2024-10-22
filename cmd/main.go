package main

import (
	"authorization_service/internal/application"
	"authorization_service/internal/core/configuration"
	"authorization_service/internal/infrastructure/repository"
	"authorization_service/internal/infrastructure/s3"
	"authorization_service/internal/infrastructure/sms"
	"authorization_service/internal/transport/amqp"
	"authorization_service/internal/transport/grpc"
	"authorization_service/internal/transport/http"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := configuration.Load()
	if err != nil {
		log.Fatal(err)
	}

	smsService := sms.New(&cfg.SMS)

	s3, err := s3.New(&cfg.S3)
	if err != nil {
		log.Fatal(err)
	}

	repository, err := repository.New(&cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	publisher, err := amqp.New(&cfg.AMQP)
	if err != nil {
		log.Fatal(err)
	}

	useCase := application.New(repository, smsService, s3, publisher)

	httpServer := http.New(useCase)
	httpServer.Register()

	go httpServer.ListenAndServe(cfg.Server.HttpSocket)

	grpcServer := grpc.New(useCase)
	go grpcServer.Run(cfg.Server.GrpcSocket)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	c, shutdown := context.WithTimeout(context.Background(), time.Second*5)
	defer shutdown()

	httpServer.Shutdown(c)
	repository.Close()
}
