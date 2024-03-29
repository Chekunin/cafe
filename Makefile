MODULE = cafe
VERSION = $(shell git describe --tags --always --dirty)
OS = $(shell uname | tr A-Z a-z)
BUILD_IMAGE ?= golang:1.17
SERVICES = $(shell ls cmd)
ABSENT_BINARIES = $(shell BINARIES="$(shell [ -d bin ] && ls bin)"; \
	echo $(SERVICES) $${BINARIES} $${BINARIES} | tr ' ' '\n' | sort | uniq -u)

.PHONY: help \
		lint \
		migrate \
		migrate/down \
		migrate/status \
		build \
		build-absent \
		clean \
		logs \
		restart \
		start \
		stop \
		test

build/%:
	@echo 'Building "$*"... '
	@mkdir -p $(CURDIR)/bin && \
	docker run \
		-i \
		--rm \
		-e SSH_AUTH_SOCK=$(SSH_AUTH_SOCK) \
		-v $(SSH_AUTH_SOCK):$(SSH_AUTH_SOCK):ro \
		-v cafe_gocache:/gocache \
		-v $(CURDIR)/cmd/$*:/opt/app/cmd \
		-v $(CURDIR)/pkg:/opt/app/pkg \
		-v $(CURDIR)/doc:/opt/app/doc \
		-v $(CURDIR)/go.mod:/opt/app/go.mod \
		-v $(CURDIR)/go.sum:/opt/app/go.sum \
		-v $(CURDIR)/configs/$*/config.yaml:/opt/app/config.yaml \
		-v $(CURDIR)/bin:/opt/app/bin \
		$(BUILD_IMAGE) \
		sh -c '\
			cd /opt/app \
			&& go env -w \
				GOMODCACHE=/gocache/mod \
				GOCACHE=/gocache/build \
			&& GIT_SSH_COMMAND="ssh -o StrictHostKeyChecking=no" \
				go build -v -o bin/$* cmd/main.go \
		' \
		&& echo 'Build for "$*" is completed'

build: $(addprefix build/, $(SERVICES)) ## Build all app binaries

build-absent: $(addprefix build/, $(ABSENT_BINARIES)) ## Build absent app binaries

docker/create-network: ## Create global docker network for app services
	@docker network create cafe || true

start: build-absent docker/create-network ## Run docker container with App
	@docker-compose up -d
	@if ! docker-compose exec db sh -c '\
				psql -t -U $${POSTGRES_USER} -d $${POSTGRES_DB} -c "SELECT 1;" \
			' | grep 1 > /dev/null; \
	then \
		echo "\n==================== Initializing DB ===================\n"; \
		$(MAKE) --no-print-directory db/seed; \
	fi

stop: ## Stop docker container with App
	@docker-compose down --remove-orphans

restart: stop start

db/seed: ## Seed DB with data
	@echo -n "Seeding DB with data... "
	@docker-compose exec db sh -c 'psql -U $${POSTGRES_USER} -d $${POSTGRES_DB} -f /app/db/seeders/dev.sql > /dev/null'
	@echo "Done!"

swagger-update: ## Update swagger docs
	@echo "##### Updating swagger docs #####"
	@echo "Updating docs for client_gateway_service..."
	@swag init --parseDependency --parseDepth 3 -g main.go -d ./cmd/client_gateway_service -o ./doc/client_gateway_service/v1