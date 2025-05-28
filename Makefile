build:
	@echo "Building frontend..."
	cd assets/frontend && bun run build

run: build
	@echo "Running Go application..."
	go run .
