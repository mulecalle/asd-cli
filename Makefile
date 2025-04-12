APP_NAME = asd

default: build

.PHONY: build
build:
	@echo "Building binary ..."
	@docker build --output=. --target=binary .
	@chmod +x $(APP_NAME)

.PHONY: install
 install: uninstall build
	@echo "Copying binary to bin folder ..."
	@sudo mv $(APP_NAME) /usr/local/bin

.PHONY: run
run:
	./$(APP_NAME)

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test: fmt
	go test ./...

.PHONY: coverage
coverage: fmt clean-coverage 
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: clean-coverage
clean-coverage:
	@echo "Removing coverage files ..."
	@rm -f coverage.out coverage.html

.PHONY: uninstall
uninstall:
	@echo "Removing binary from root folder ..."
	@rm -rf ./$(APP_NAME) > /dev/null
	@echo "Removing binary from bin folder ..."
	@sudo rm -rf /usr/local/bin/$(APP_NAME)
