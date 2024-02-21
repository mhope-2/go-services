
build:
	docker-compose build

up:
	@make build
	docker-compose -f docker-compose.yml up

logs:
	docker-compose logs -f

rm:
	docker-compose rm  -sfv

start:
	docker-compose start

stop:
	docker-compose stop

test:
	go test -v -cover ./...
