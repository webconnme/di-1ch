package main

import (
	"github.com/webconnme/go-webconn"
	"github.com/webconnme/go-webconn-gpio"
	"log"
)

var client webconn.Webconn
var g *gpio.Gpio

func D1_IN(buf []byte) error{

	data := string(buf)
	log.Println(">>>In data : ",data)

	din, err := g.In();
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(">>>di data : ",din)

	return nil
}

func main() {

	g = &gpio.Gpio{248, gpio.IN}
	err := g.Open()
	if err != nil {
		log.Println(err)
	}
	defer g.Close()

	client = webconn.NewClient("http://192.168.4.180:3004/v01/di1ch/80")
	client.AddHandler("request",D1_IN)

	client.Run()
}
