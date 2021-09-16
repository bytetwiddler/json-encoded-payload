
UNAME := $(shell uname | tr [:upper:] [:lower:])

os: 
	make $(UNAME)

all: darwin linux windows

darwin:
	GOOS=darwin GOARCH=amd64 go build -o 'example-darwin-amd64' main.go
	GOOS=darwin GOARCH=arm64 go build -o 'example-darwin-arm64' main.go

linux:
	GOOS=linux GOARCH=amd64 go build -o 'example-linux-amd64' main.go
	GOOS=linux GOARCH=386 go build -o 'example-linux-386' main.go

windows:
	GOOS=windows GOARCH=amd64 go build -o 'example-windows-amd64.exe' main.go
	GOOS=windows GOARCH=386 go build -o 'example-windows-386.exe' main.go

run:
	go run main.go

clean:
	rm -f example-* *.exe new_* *.zip

zips: all
	zip windows.zip example-windows-*exe
	zip linux.zip example-linux-*
	zip mac.zip example-darwin-*

