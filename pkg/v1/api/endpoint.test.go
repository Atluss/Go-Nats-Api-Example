package api

import (
	"encoding/json"
	"fmt"
	"github.com/Atluss/Go-Nats-Api-Example/pkg/v1"
	"github.com/Atluss/Go-Nats-Api-Example/pkg/v1/config"
	"github.com/gorilla/mux"
	"github.com/nats-io/go-nats"
	"log"
	"net/http"
	"sync"
	"time"
)

// User base struct for users
type User struct {
	Id   string
	Name string
}

func (u User) String() string {
	return fmt.Sprintf("Id: %s, Name: %s", u.Id, u.Name)
}

// NewEndPointV1Test constructor for /v1/test/{id}
func NewEndPointV1Test(set *config.Setup) (*v1test, error) {

	url := fmt.Sprintf("/%s/test/{id}", V1ApiQueue)
	if err := CheckEndPoint(V1ApiQueue, url); err != nil {
		return nil, err
	}
	endP := v1test{
		Setup: set,
		Url:   url,
	}
	err := endP.SetRouteAndNats()
	AddEndPoint(V1ApiQueue, endP.Url)
	return &endP, err
}

type v1test struct {
	Setup *config.Setup
	Url   string
}

// SetRouteAndNats setup route link and Nats queue
func (obj *v1test) SetRouteAndNats() error {
	// register queue for API and url
	_, err := obj.Setup.Nats.QueueSubscribe(obj.Url, V1ApiQueue, obj.NatsQueue)
	if !v1.LogOnError(err, "Can't listen nats queue") {
		return err
	}
	// register frontend url
	obj.Setup.Route.HandleFunc(obj.Url, obj.Request)
	return nil
}

// Request setup mux answer
func (obj *v1test) Request(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	myUser := User{Id: vars["id"]}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {

		data, err := json.Marshal(&myUser)
		if err != nil || len(myUser.Id) == 0 {
			v1.LogOnError(err, "warning: Problem with parsing the user Id: %s")
			w.WriteHeader(500)
			return
		}

		msg, err := obj.Setup.Nats.Request(obj.Url, data, 100*time.Millisecond)
		if err == nil && msg != nil {

			myUserWithName := User{}

			err := json.Unmarshal(msg.Data, &myUserWithName)
			log.Printf("Request user: %s", myUserWithName)

			if err == nil {
				myUser = myUserWithName
			}

			w.Header().Set("Content-Type", "application/json")

			v1.LogOnError(json.NewEncoder(w).Encode(myUserWithName), "warning")
		}
		wg.Done()

	}()
	wg.Wait()

}

// NatsQueue add new queue
func (obj *v1test) NatsQueue(m *nats.Msg) {

	users := map[string]string{
		"1": "Ilya",
		"2": "Yana",
		"3": "Olga",
		"4": "Shokin",
	}

	myUser := User{}
	err := json.Unmarshal(m.Data, &myUser)

	if err != nil {
		log.Println(err)
		return
	}

	myUser.Name = users[myUser.Id]

	data, err := json.Marshal(&myUser)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Replying to ", m.Reply)

	err = obj.Setup.Nats.Publish(m.Reply, data)
	v1.LogOnError(err, "warning")
}
