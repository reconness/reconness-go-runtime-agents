package main

import (
	"time"
	"log"
	"math/rand"
	"fmt"
	"github.com/kardianos/service"
//	"github.com/streadway/amqp"
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
	
	//conn, err := amqp.Dial("amqp://reconness:reconness@rabbitmq:5672/")
	//if err != nil {
	//	fmt.Println(err)
	//	panic(1)
	//}
	//defer conn.Close()
	
	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	
	// Do work here
	// infinite print loop
        for {
          fmt.Println("Hello, World!")

          // wait random number of milliseconds
          Nsecs := rand.Intn(3000)
          time.Sleep(time.Millisecond * time.Duration(Nsecs))
        }	
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
