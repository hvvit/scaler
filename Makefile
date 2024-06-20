
clean:
	rm scaler

run:
	go run ./... watch

build:
	CGO_ENABLED=0
	go build -o scaler

build-linux:
	CGO_ENABLED=0
	GOOS=linux
	GOARCH=amd64
	go build -o scaler

build-mac:
	CGO_ENABLED=0
	GOOS=darwin
	GOARCH=amd64
	go build -o scaler

build-windows:
	CGO_ENABLED=0
	GOOS=windows
	GOARCH=amd64
	go build -o scaler

clean-build:
	rm scaler
	CGO_ENABLED=0
	go build -o scaler

test:
	go test -v ./...