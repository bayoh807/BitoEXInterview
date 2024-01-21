TARGETS = tinder-Server

CMD_NAME = "tinder-Server"

init_for_docker:
	make init build $(CMD_NAME)
init:
	rm -f go.mod
	rm -f go.sum
	go mod init $(TARGETS)
	go mod tidy

build: $(TARGETS)

run:
	go run main.go

clean:
	go clean
	rm -rf build/*


.PHONY: build

