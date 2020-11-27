run-dev:
	SERVER_ADDR=localhost:8000 go run cmd/server/main.go

build:
	 @make clean
	CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=6 go build -o deploy/assets/server cmd/server/main.go

deploys:
	ansible-playbook -i deploy/inventories/production/hosts deploy/deploy.yml

tests:
	go test ./...

clean:
	go fmt ./...