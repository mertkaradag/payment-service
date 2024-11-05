run:
	@echo "Running the project"
	go run ./cmd/payment/main.go

unit-test:
	@echo "Starting unit test"
	go clean -testcache && go test  ./... -v -short

code-coverage-test:
	@echo "Starting code covarege"
	go test -tags=unit -count=1 -covermode=atomic -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html