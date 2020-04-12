run-dev:
	SERVER_ADDR=localhost:8000 go run cmd/server/main.go

build:
	cd cmd/server && CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=6 go build

deploy:
	@make build && scp server rainSensor