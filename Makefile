
docker-logs:
	docker logs -f rain_sensor

tools-ci: tool-golangci-lint tool-fumpt
tools-local: tool-golangci-lint tool-moq tool-fumpt	 tool-jsonschema tool-json-lint

tool-golangci-lint:
	devops/scripts/goget.sh github.com/golangci/golangci-lint/cmd/golangci-lint@v1.47.3

tool-fumpt:
	devops/scripts/goget.sh mvdan.cc/gofumpt

tool-moq:
	devops/scripts/goget.sh github.com/matryer/moq

tool-jsonschema:
	devops/scripts/goget.sh github.com/atombender/go-jsonschema/...
	devops/scripts/goget.sh github.com/atombender/go-jsonschema/cmd/gojsonschema

tool-json-lint:
	go get github.com/santhosh-tekuri/jsonschema/cmd/jv

test:
	go test -race ./...

test-functional:
	go test -tags functional -race ./functional_test/... --count=1

clean:
	go fmt ./...

lint:
	golangci-lint run
	go mod tidy -v && git --no-pager diff --quiet go.mod go.sum

import-jsonschema:
	devops/scripts/import_jsonschema.sh

json-lint:
	devops/scripts/json-lint.sh

build:
	 @make clean
	CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=6 go build -o devops/ansible/assets/server cmd/server/main.go

deploy:
	ansible-playbook -i devops/ansible/inventories/production/hosts devops/ansible/deploy.yml

docker-up:
	docker-compose up -d --build rain_sensor

docker-down:
	docker-compose down

docker-ps:
	docker-compose ps

docker-exec:
	docker exec -it rain_sensor sh