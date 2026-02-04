test:
	docker compose exec api go test ./internal/config -v
lint:
	docker run --rm -t -v ${CURDIR}:/app -w /app golangci/golangci-lint:v2.8.0 golangci-lint run
start:
	docker compose up -d
ci: start test lint