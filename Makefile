ARTIFACTS=${PWD}/bin

.PHONY: build
build: main.go
	GOARCH=amd64 GOOS=linux go build -o ${ARTIFACTS}/composer-update.am64.linux
	GOARCH=amd64 GOOS=darwin go build -o ${ARTIFACTS}/composer-update.amd64.darwin