# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Function to check for required environment variables
check-env-vars:
	@if [ -z "$(APP_NAME)" ]; then \
		echo "Error: APP_NAME is not set"; \
		exit 1; \
	fi

	@if [ -z "$(PORT)" ]; then \
		echo "Error: PORT is not set"; \
		exit 1; \
	fi

	@if [ -z "$(MIGRATIONS_DIR)" ]; then \
		echo "Error: MIGRATIONS_DIR is not set"; \
		exit 1; \
	fi
	
	@if [ -z "$(DB_DRIVER)" ]; then \
		echo "Error: DB_DRIVER is not set"; \
		exit 1; \
	fi
	@if [ -z "$(DSN)" ]; then \
		echo "Error: DSN is not set"; \
		exit 1; \
	fi

	@if [ -z "$(SQLC_OUT)" ]; then \
		echo "Error: SQLC_OUT is not set"; \
		exit 1; \
	fi

all: check-env-vars clean migrate-up compile-sql compile-proto build run

clean: check-env-vars
	rm -rf $(SQLC_OUT)/*

migrate-up: check-env-vars
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DSN)" up

migrate-down: check-env-vars
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DSN)" down

compile-sql: check-env-vars
	rm -rf $(SQLC_OUT)/* && sqlc generate

build: check-env-vars
	go build -o bin/$(APP_NAME) cmd/api/main.go

run: check-env-vars
	bin/$(APP_NAME)
