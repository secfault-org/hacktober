build:
	GOOS=linux go build -ldflags "-w -s" -o hacktober main.go

run:
	go run main.go

all: build
