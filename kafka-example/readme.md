# Kafka Example
This is an example project to go along with our engineering blog titled
"Getting Started with Kafka Using Go."  If you came here from that article,
then you are in the right spot.  If not then you may still be anyway.

This project is meant to provide a simple example on how to use Kafka with Go.

It is a simple webservice designed to return and build scores for a users 
particular skill set

#### Getting up and going
If you have Docker already installed then you all you need to do is
run 

```$ docker-compose up```

from the terminal and the correct container will download and startup
configured correctly.

You may install Kafka and Zookeeper locally if you wish, and there are guides
out there on how to do that.

I highly suggest going the docker route.

#### How to run
The easiest way to run the project is to build and then run it.

```bash
$ go build && ./kafka-example
```

This will start the app and it will be listening on port 3000

#### How to call the service

I personally used Postman to test the service, but feel free to use what ever
you prefer.

Posting and object to ```http://localhost:3000/api/skills``` with this shape
```
{ 
    "id":"user1", 
    "skills":[
        "Services", 
        "ReactJS", 
        "Kafka", 
        "Golang"
    ] 
}
```

should produce a 202 response and message should start scrolling in the terminal
window

Issuing a get request to 

```
http://localhost:3000/api/skills?userID=user1&skill=ReactJS
```

should return an object like

```
{
  "skill_name": "ReactJS",
  "score": 6.260835,
  "last_scored": "2017-01-04T20:18:44.107199795-05:00"
}
```