# GO SERVICES
This is an attempt to setup 3 go services and a RabbitMQ connection in a docker environment.

### User Service
The User Service is a simple service that returns a user object based on specified user IDs. The service also simulates 
some delays and faults that should be handled in the order service.

### Product Service  
The Product service is another simple service with three fixed products that are retrieved by their product codes. This
service also some delays and faults

### Order Service  
The Order Service is a more extensive service the receives a POST request containing a user ID and product code, fetches the 
corresponding data from the user and product services, saves the order, and then publishes a message to RabbitMQ.
Sample create order request body:
```json
{
    "user_id": "e6f24d7d1c7e",
    "product_code": "product1"
}
```

## Setup Guide

### Requirements

- Docker

### Instructions

1. Clone the repo and in each service, make a copy of .env.sample as .env and update the environment variables as you deem fit

```bash
git clone https://github.com/mhope-2/go-services.git
cp .env.sample .env
```

2. Start the services in a docker environment

```bash
docker-compose up
```
or
```bash
make up
```

### Running Tests
For the user and product services:
```bash
go test -v -cover ./...
```

For the order service, since the dependencies (Postgres and RabbitMQ) aren't completely isolated in the tests as disposable 
docker containers, you'll have to keep them up before running the test command