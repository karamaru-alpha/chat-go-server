ENV_LOCAL_FILE = env.local
ENV_LOCAL = $(shell cat $(ENV_LOCAL_FILE))

.PHONY:run
run:
	$(ENV_LOCAL) docker-compose up

.PHONY:reset_migration
reset_migration:
	$(ENV_LOCAL) sh ./docker/mysql/db/init/init-mysql.sh

.PHONY:test
test:
	$(ENV_LOCAL) go test ./test/...

.PHONY:lint
lint:
	go mod tidy
	golangci-lint run ./...
