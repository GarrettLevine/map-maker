version=0.1.0


all:
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@echo "  build       - build the dist binary"
	@echo "  clean       - clean the dist build"
	@echo "  deps        - update dependencies"
	@echo "  test        - standard go test"

build: clean
	@go vet $$(go list ./... | grep -v /vendor/)
	@go build -o ./bin/map_maker-$(version).bin -ldflags "-X main.version=$(version)" cmd/main.go

clean:
	@rm -rf ./bin
	@mkdir ./bin

deps:
	@glide up

test:
	@go test $$(go list ./... | grep -v /vendor/)