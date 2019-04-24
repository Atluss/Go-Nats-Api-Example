Go Nats API Example
=====================

This example show how to use [Nats](https://www.nats.io/) to create API.
Project use [Go modules](https://github.com/golang/go/wiki/Modules) to download dependencies.
You can see dependencies in `go.mod`

Docker
-----------------------------------
How to install: 
 1. [Install Docker-CE (ubuntu)](https://docs.docker.com/install/linux/docker-ce/ubuntu/)
 2. [Install Docker compose](https://docs.docker.com/compose/install/)
 4. In project root: `sudo docker-compose up`
 
 Settings file
 -----------------------------------
 API settings file use [RFC7159](https://tools.ietf.org/html/rfc7159)
 
 Example settings.json (all settings is required) :
 ```json
 {
   "name": "api",
   "version": "1.0.0",
   "nats": {
     "version" : "1.4.2",
     "reconnectedWait" : 5,
     "address" : [
       {
         "host" : "localhost",
         "port" : "54222"
       },
       {
         "host" : "localhost",
         "port" : "54222"
       }
     ]
   }
 }
 ```