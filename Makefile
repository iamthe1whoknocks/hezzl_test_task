## build app
service: 
	go mod tidy && go mod download && go run ./cmd/app

## run docker-compose
up:
	docker-compose up --build -d && docker-compose logs -f

down:
	docker-compose down --remove-orphans

logs:
	docker-compose logs -f

migrations:
	migrate create -ext sql -dir db/migrations -seq create_users_table