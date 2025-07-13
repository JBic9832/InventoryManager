build:
	@go build -o bin/inventory-tracker

run: build
	@./bin/inventory-tracker
