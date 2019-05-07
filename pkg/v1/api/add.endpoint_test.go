package api

import (
	"log"
	"testing"
)

func TestAddEndPoint(t *testing.T) {

	url := "/test"

	AddEndPoint(V1ApiQueue, url)

	for i, k := range endPoints {
		log.Println("=============")
		log.Printf("Queue: %s", i)

		for _, nc := range k {
			log.Printf("%s", nc.url)
		}
		log.Println("=============")
	}
}
