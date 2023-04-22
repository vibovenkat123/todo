BIN="./bin"

.PHONY: fmt lint test install_deps clean

default: all

all: clean fmt lint align_structs test install_deps build

build: install_deps
	$(info ******************** building the app ********************)
	mkdir -p $(BIN)
	go build -o $(BIN)/todo main.go

fmt:
	$(info ******************** formatting code ********************)
	gofmt -w -s -l .

lint:
	$(info ******************** linting ********************)
	golangci-lint run -v

align_structs:
	$(info ******************** aligning struct fields ********************)
	fieldalignment -fix ./...

test: install_deps
	$(info ******************** running tests ********************)
	go test -v ./...

install_deps:
	$(info ******************** downloading dependencies ********************)
	go get -v ./...

clean:
	$(info ******************** cleaning bin files ********************)
	rm -rf $(BIN)
