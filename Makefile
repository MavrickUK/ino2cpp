LDFLAGS = -ldflags "-s -w"
APPNAME = ino2cpp
BUILD = "0.2"

.DEFAULT_GOAL := run

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build -o bin/${APPNAME}.exe main.go
.PHONY:build

tiny:
	env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o bin/${APPNAME}_tiny.exe main.go

upx: tiny
	upx -9 -o bin/${APPNAME}_Packed.exe bin/${APPNAME}_tiny.exe

comwin32:
	env GOOS=windows GOARCH=386 go build ${LDFLAGS} -o bin/${APPNAME}-windows-386.exe main.go

comwin64:
	env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o bin/${APPNAME}-windows-amd64.exe main.go

linux:
	env GOOS=linux GOARCH=amd64 go build -o bin/${APPNAME}-linux-amd64 main.go

run: build
	.\bin\${APPNAME}.exe
	
update:
	go get -u all

compile:
	# 64-Bit
	#
# FreeBDS
# env GOOS=freebsd GOARCH=amd64 go build -o bin/${APPNAME}-${BUILD}-freebsd-amd64 main.go
	# MacOS
	env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o bin/${APPNAME}-darwin-amd64 main.go
	# Linux
	env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o bin/${APPNAME}-linux-amd64 main.go
	upx -9 -o bin/${APPNAME}-linux-amd64_Packed bin/${APPNAME}-linux-amd64
	# Windows
	env GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o bin/${APPNAME}-windows-amd64.exe main.go
	upx -9 -o bin/${APPNAME}-windows-amd64_Packed.exe bin/${APPNAME}-windows-amd64.exe