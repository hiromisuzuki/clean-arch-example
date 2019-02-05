PKG_LIST = $(shell go list ./... | grep -v /vendor/)

local: init dep fresh
build: stop start

.PHONY: help
help: ## print this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: init
init: ## initialize
	go get github.com/golang/dep/cmd/dep && \
  go get github.com/go-sql-driver/mysql && \
  go get github.com/swaggo/swag/cmd/swag && \
  go get bitbucket.org/liamstask/goose/cmd/goose && \
  go get github.com/lib/pq && \
  go get github.com/pilu/fresh
	swag init

.PHONY: install
install: ## install into $GOPATH/bin
	go install -v

.PHONY: run
run:
	go run main.go

.PHONY: fresh
fresh:
	fresh main.go

.PHONY: migrate
migrate: ## migrate database for docker-compose
	docker-compose exec api goose up

.PHONY: start
start: ## start docker-compose
	docker-compose up --build

.PHONY: stop
stop: ## stop docker-compose
	docker-compose down

.PHONY: log
log: ## logging services on docker-compose
	docker-compose logs -f

.PHONY: apilog
apilog: ## logging api on docker-compose
	docker-compose logs -f api

.PHONY: dblog
dblog: ## logging dblog on docker-compose
	docker-compose logs -f dblog

.PHONY: test
test: vet lint ## run tests
	go test $(PKG_LIST)

.PHONY: race
race: vet lint ## run tests with racedetector
	go test -race $(PKG_LIST)

.PHONY: vet
vet: ## vet sources
	go vet $(PKG_LIST)

.PHONY: lint
lint: $(GOPATH)/bin/golint ## lint sources
	golint -set_exit_status -min_confidence=0.4 $(PKG_LIST)

.PHONY: doc
doc: ## run godoc server on http://localhost:6060/pkg
	godoc -http=":6060"

.PHONY: dep
dep: ## ensure dependencies are met
	dep ensure -v

.PHONY: depup
depup: ## update dependencies
	dep ensure -v -update $(PKG) && dep prune && find ./vendor -name '*_test.go' -delete

.PHONY: swag
swag: 
	docker-compose exec api swag init
