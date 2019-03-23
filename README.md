Go Nats API Example
=====================

This example show how to use [Nats](https://www.nats.io/) to create API.
Project use [Go modules](https://github.com/golang/go/wiki/Modules) to download dependencies.

Docker
-----------------------------------
How to install: 
 1. [Install Docker-CE (ubuntu)](https://docs.docker.com/install/linux/docker-ce/ubuntu/)
 2. [Install Docker compose](https://docs.docker.com/compose/install/)
 3. Unzip docker/docker.zip to folder(Nats 1.4.1, Postgres 11.2)
 4. In this folder: `sudo docker-compose up`
 
 Setting file
 -----------------------------------
 API settings file use [RFC7159](https://tools.ietf.org/html/rfc7159)
 
 Example settings.json (all settings if required) :
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