build:
	@go build -v -o bin/inventory-tracker 

run: build
	@./bin/inventory-tracker
