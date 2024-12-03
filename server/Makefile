# Define application name
APP_NAME=talon

# Define app mode (default to dev if not set)
APP_MODE ?= dev

build:
	set APP_MODE=$(APP_MODE) && go build -o $(APP_NAME) .

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

# Run the app in development mode
dev:
	@echo "Running in development mode..."
ifeq ($(OS),Windows_NT)
	set APP_MODE=dev && air
else
	APP_MODE=dev air
endif

# Clean build files
clean:
	rm -f $(APP_NAME)
