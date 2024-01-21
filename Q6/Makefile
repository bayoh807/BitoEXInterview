TARGETS = tinder-server

CMD_NAME = "tinder-server"

init_for_docker:
	make init build
init:
	rm -f go.mod
	rm -f go.sum
	go mod init $(TARGETS)
	go mod tidy

build:
	go build -o build/main main.go

run:
	go run main.go

clean:
	go clean
	rm -rf build/*


.PHONY: build

