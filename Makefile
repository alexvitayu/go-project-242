# Сборка проекта
build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

# Запуск линтера
run:
	golangci-lint run

# Запуск линтера
test:
	go test -v ./...