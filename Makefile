## build app
service: 
	go mod tidy && go mod download && go run ./cmd/app

## run docker-compose
up:
	docker-compose up --build -d && docker-compose logs -f

## stop services
down:
	docker-compose down --remove-orphans

## watch logs
logs:
	docker-compose logs -f