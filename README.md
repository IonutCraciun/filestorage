## Golang webservice for file operations (create, update, read, delete)

Endpoints:

1. /
2. /view
3. /update
4. /newfile
5. /delete

Port: 8080

Command to start server: 
```
go run cmd/main/main.go
```

Command to start the consumers:
```
cd consumer
go run consumer.go   // for a queue that listens to messages with routing key "file.update"
go run consumer.go file.create  // for a queue that listens to messages with routing key "file.create"
```


Command to start a rabbitmq instance using docker
```
 docker run --detach \
 --name rabbitmq \
 -p 5672:5672 \
 -p 15672:15672 \
 rabbitmq:3-management
```