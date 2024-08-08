APP=dockit

build-local:
	go build -o ${APP} main.go

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${APP} main.go

build-macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${APP} main.go

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${APP}.exe main.go

clean:
	go clean

