# Define application name
APP_NAME=talon

# Build the application binary
build:
	go build -o $(APP_NAME) .

# Apply all migrations
migrate-up: build
	./$(APP_NAME) migrate up

# Rollback the last migration
migrate-down: build
	./$(APP_NAME) migrate down

# Check migration status
migrate-status: build
	./$(APP_NAME) migrate status

# Start the application
run: build
	./$(APP_NAME)

dev:
	air

# Clean build files
clean:
	rm -f $(APP_NAME)
