# Args
day = 1

# Default task
all: run

# Run the application
run:
	go run ./src/main.go ${day}

# Run tests
test:
	go test main_test.go
