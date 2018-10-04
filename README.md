## Demo Produce and Consume Message using Golang and Apache Kafka

#### TODO
- Consumer connect using Zookeeper

### Getting started

- Start Kafka and Zookeeper
```shell
$ docker-compose up
```

- Start Producer

  install dependencies
```shell
$ glide install
```
```shell
$ go build
```
```shell
$ ./producer
```

- Send Message to Kafka topic
```shell
curl -X POST \
http://localhost:3000/api/send \
-H 'content-type: application/json' \
-d '{
    "from": "Wuriyanto",
        "content":{
            "header": "This is Message 2",
            "body": "Hello Kafka"
        }
    }'
```

- Start Consumer

  install dependencies
```shell
$ glide install
```
```shell
$ go build
```
```shell
$ ./consumer
```