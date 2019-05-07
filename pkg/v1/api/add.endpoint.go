package api

import (
	"fmt"
)

// V1ApiQueue api version
const V1ApiQueue = "v1"

type endPoint struct {
	url string
}

var endPoints map[string][]endPoint

// AddEndPoint add endpoint to reg endpoints array
func AddEndPoint(queue, url string) {

	if endPoints == nil {
		endPoints = map[string][]endPoint{}
	}

	_, ok := endPoints[queue]
	if !ok {
		endPoints[queue] = []endPoint{}
	}

	endPoints[queue] = append(endPoints[queue], endPoint{url: url})
}

// CheckEndPoint endpoint in reg endpoints
func CheckEndPoint(queue, url string) error {

	if endPoints == nil {
		return nil
	}

	_, ok := endPoints[queue]
	if !ok {
		return nil
	}

	for _, nc := range endPoints[queue] {
		if nc.url == url {
			return fmt.Errorf("endpoint: %s already set", url)
		}
	}

	return nil
}
