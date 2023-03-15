LDFLAGS = -ldflags "-s -w"
APPNAME = ino2cpp

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

run: build
	.\bin\${APPNAME}.exe
	
update:
	go get -u all

compile:
	# 32-Bit Systems
	# FreeBDS
	env GOOS=freebsd GOARCH=386 go build -o bin/${APPNAME}-freebsd-386 main.go
	# MacOS
	env GOOS=darwin GOARCH=386 go build -o bin/${APPNAME}-darwin-386 main.go
	# Linux
	env GOOS=linux GOARCH=386 go build -o bin/${APPNAME}-linux-386 main.go
	# Windows
	env GOOS=windows GOARCH=386 go build -o bin/${APPNAME}-windows-386.exe main.go
# 64-Bit
	# FreeBDS
	env GOOS=freebsd GOARCH=amd64 go build -o bin/${APPNAME}-freebsd-amd64 main.go
	# MacOS
	env GOOS=darwin GOARCH=amd64 go build -o bin/${APPNAME}-darwin-amd64 main.go
	# Linux
	env GOOS=linux GOARCH=amd64 go build -o bin/${APPNAME}-linux-amd64 main.go
	# Windows
	env GOOS=windows GOARCH=amd64 go build -o bin/${APPNAME}-windows-amd64.exe main.go
