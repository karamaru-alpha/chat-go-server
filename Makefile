ENV_LOCAL_FILE = env.local
ENV_LOCAL = $(shell cat $(ENV_LOCAL_FILE))

run:
	$(ENV_LOCAL) docker-compose up

reset_migration:
	$(ENV_LOCAL) sh ./docker/mysql/db/init/init-mysql.sh

lint:
	go mod tidy
	golangci-lint run ./...
