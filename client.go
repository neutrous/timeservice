// -*- mode: go -*-

// This package defines the client side to receive the unified time
// datas from the server.

package tu

import (
	"log"
	
	"github.com/neutrous/notify/entity"
)

type Client struct {
	env entity.CommEnv
	sub entity.Subscriber
}

// NewClient creates a new instance of client.
func NewClient(address string) *Client {
	client := Client{}
	client.sub.AppendAddress(entity.Address(address))
	
	env, err := entity.CreateZMQCommEnv(true)
	if err != nil {
		return nil
	}
	if client.sub.InitialConnecting(env) != nil {
		return nil
	}
	return &client
}

// Close releases the instance's resource.
func (obj *Client) Close() {
	if obj != nil {
		obj.env.Close()
	}
}

// Subscribe tell the underly implementation to receive the time datas.
func (obj *Client) Subscribe() {
	handler := &TimeDeserializer{}
	if err := obj.sub.Subscribe(handler); err != nil {
		log.Fatalln("Subscribe failure.")
		return
	}
	log.Println("Start receiving time datas.")
	obj.sub.ReceivingEvent()
}



















