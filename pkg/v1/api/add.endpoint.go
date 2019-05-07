package api

import (
	"fmt"
)

// V1ApiQueue api version
const V1ApiQueue = "v1"

type endPoint struct {
	url string
}

var EndPoints map[string][]endPoint

// AddEndPoint add endpoint to reg endpoints array
func AddEndPoint(queue, url string) {

	if EndPoints == nil {
		EndPoints = map[string][]endPoint{}
	}

	_, ok := EndPoints[queue]
	if !ok {
		EndPoints[queue] = []endPoint{}
	}

	EndPoints[queue] = append(EndPoints[queue], endPoint{url: url})
}

// CheckEndPoint endpoint in reg endpoints
func CheckEndPoint(queue, url string) error {

	if EndPoints == nil {
		return nil
	}

	_, ok := EndPoints[queue]
	if !ok {
		return nil
	}

	for _, nc := range EndPoints[queue] {
		if nc.url == url {
			return fmt.Errorf("endpoint: %s already set", url)
		}
	}

	return nil
}
