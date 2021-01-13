package main

import (
	"time"
	"log"
	"fmt"
	"github.com/kardianos/service"
	"github.com/streadway/amqp"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) run() {
	fmt.Println("Go RabbitMQ Tutorial")
	
  	time.Sleep(time.Millisecond * time.Duration(30000))
          
	conn, err := amqp.Dial("amqp://reconness:reconness@rabbitmq:5672/")
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
	
	
	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	
	// Let's start by opening a channel to our RabbitMQ instance
    	// over the connection we have already established
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()
	
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	
	// Handle any errors if we were unable to create the queue
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(q)
	
    	msgs, err := ch.Consume(
		"hello",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
    	
    	defer conn.Close()
    	
	
    	// Handle any errors if we were unable to create the queue
	if err != nil {
		fmt.Println(err)
	}
	
	forever := make(chan bool)
	go func() {
		for d := range msgs {
		
			fmt.Printf("Recieved Message: %s\n", d.Body)
			
			err = ch.Publish(
				"",
				"hello",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte("64 bytes from w2.src.vip.bf1.yahoo.com (74.6.136.150): icmp_seq=1 ttl=51 time=77.9 ms"),
				},
			)

			if err != nil {
				fmt.Println(err)
			}
			
		    	fmt.Println("Successfully Published Message to Queue")
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "GoServiceTest",
		DisplayName: "Go Service Test",
		Description: "This is a test Go service.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
