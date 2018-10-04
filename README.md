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