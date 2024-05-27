APP_NAME = asd

default: build

.PHONY: build
build: clean
	@echo "Building binary ..."
	@docker build --output=. --target=binary .
	@chmod +x $(APP_NAME)

.PHONY: publish
 publish:
	@echo "Copying binary to bin folder ..."
	@cp ./$(APP_NAME) /usr/local/bin

.PHONY: run
run:
	./$(APP_NAME)

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test: fmt
	go test ./...

.PHONY: clean
clean:
	@echo "Removing binary from root folder ..."
	@rm -rf ./$(APP_NAME) > /dev/null
	@echo "Removing binary from bin folder ..."
	@rm -rf /usr/local/bin/$(APP_NAME)
