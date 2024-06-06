
## run docker-compose
up:
	docker-compose up --build -d && docker-compose logs -f

## stop services
down:
	docker-compose down --remove-orphans

## watch logs
logs:
	docker-compose logs -f

### check by golangci linter
linter-golangci: 
	golangci-lint run --config=.golangci.yml ./...

mock:
	mockgen -source ./internal/usecase/interfaces.go -package usecase_test > ./internal/usecase/mock_test.go

unit-test:
	go test -v -cover -race ./internal/...