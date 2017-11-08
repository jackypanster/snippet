package main

import (
	"github.com/parnurzeal/gorequest"
	"log"
	"net/http"
)

func main() {
	request := gorequest.New()
	resp, body, errs := request.Get("http://127.0.0.1:8080/").End()
	if errs != nil {
		log.Printf("error: %+v", errs)
	} else if resp.StatusCode != http.StatusOK {
		log.Printf("%s, %s", resp.Status, body)
	} else {
		log.Printf("%s, %s", resp.Status, body)
	}
}
