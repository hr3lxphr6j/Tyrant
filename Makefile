.PHONY: build
build:
	@./src/hack/build.sh

.PHONY: linux-build
linux-build:
	@docker run --rm -v '${PWD}:/tmp/Tyrant' golang /tmp/Tyrant/src/hack/build.sh /tmp/Tyrant

.PHONY: clean
clean:
	@rm -rf bin/

.PHONY: generate
generate:
	@go generate ./src/...