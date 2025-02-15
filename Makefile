ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST = postgresql://user:5432@localhost:5430/Avito_Merch
endif
LOCAL_BIN:=$(CURDIR)/bin
MIGRATION_FOLDER = $(CURDIR)/migrations

.all-deps: .bin-deps .add-deps
.add-deps:
	$(info Installing dependencies...)
	go get -u github.com/jackc/pgx/v5/pgxpool
	go get -u go.uber.org/zap
	go get -u github.com/redis/go-redis/v9
	go get -u golang.org/x/crypto/bcrypt
	go get -u github.com/golang-jwt/jwt/v5


.bin-deps:
	$(info Installing binary dependencies...)
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create postgres sql

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: migration-down
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down


build-compose:
	docker-compose build

compose-up-postgres:
	docker-compose up -d postgres

compose-down:
	docker-compose down

linter: 
	golangci-lint run
