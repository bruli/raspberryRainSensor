
define help
Usage: make <command>
Commands:
   help:                      Show this help information
   tool-jsonschema:           Install gojsonschema tool
   test:                      Run unit tests
   test-functional:           Run functional tests
   docker-up:                 Start docker containers
   docker-down:               Stop docker containers
   docker-ps:                 To watch all docker containers
   docker-exec                To entry into water system container
   lint:                      Execute go linter
   clean:                     To clean code
   fumpt:					  Format code
   import-jsonschema:         Import and generate DTOS from json schemas
   build:                     Compile the project
   docker-exec-builder:       Start builder docker container and entry inside it. Build project here.
   deploy:                    Deploy the code to raspberry
endef
export help

.PHONY: help
help:
	@echo "$$help"

docker-logs:
	docker logs -f rain_sensor

test:
	go test -race ./...

test-functional:
	go test -tags functional -race ./functional_test/... --count=1

.PHONY: tool-jsonschema
tool-jsonschema:
	go get github.com/atombender/go-jsonschema/...
	go install github.com/atombender/go-jsonschema@latest

clean:
	go fmt ./...

.PHONY: fumpt
fumpt:
	go tool gofumpt -w -l .

lint:
	go tool golangci-lint run

import-jsonschema:
	devops/scripts/import_jsonschema.sh

json-lint:
	devops/scripts/json-lint.sh

build:
	 @make clean
	CC=arm-linux-gnueabi-gcc CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o devops/ansible/assets/server cmd/server/main.go

deploy:
	ansible-playbook -i devops/ansible/inventories/production/hosts devops/ansible/deploy.yml

docker-up:
	docker compose up -d --build rain_sensor

docker-down:
	docker compose down

docker-ps:
	docker compose ps

docker-exec:
	docker exec -it rain_sensor bash

docker-exec-builder:
	docker build -t builder .
	docker run -it --rm -v $(shell pwd):/app builder bash