APP=dockit

build-local:
	go build -o ${APP} main.go


build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${APP} main.go

clean:
	go clean