package main

import (
	"github.com/parnurzeal/gorequest"
	"log"
)

func main() {
	request := gorequest.New()
	resp, body, errs := request.Get("http://127.0.0.1:8080/pingg").End()
	if errs != nil {
		log.Printf("%+v, %s, error: %+v", resp, body, errs)
	} else {
		log.Printf("%+v, %s, %+v", resp.Status, body, errs)
	}
}
