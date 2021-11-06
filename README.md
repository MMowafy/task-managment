## How to use

```
## To build and start app

docker-compose up --build 

- This command will pull, build image and run container for the following with dependency on each other (will take some time for services to be up and running):  
    - Postgres
    - API service for User and Task management
    
## To Run tests
 
 cd api
 go test -v ./...
 
 ### To hit APIS 

 Import postman collection Creattive AT.postman_collection.json

- Enhancment to be done:
    - Notifcation and Mailer each should be a separate microservice
    - Add message queue like rabbitMQ or Kafka for asyncronus commuication between microservices specially after creating task to send notifiction and sending ad email
    - Should be dockerized in a better way 
    - Write Integration test
    - Better dependency injection and use more interfaces in infrastructure


```

