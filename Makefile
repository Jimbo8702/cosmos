BINARY_NAME=myApp

build:
	@go mod vendor
	@echo "Building it..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "...It lives"

run: build
	@echo "Scaring it..."
	@./tmp/${BINARY_NAME} &
	@echo "...Its running"

clean:
	@echo "Spring cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

docup:
	docker-compose up -d

docDown: 
	docker-compose down

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Knocking it out..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "...it sleeps"

restart: stop start