# Собирает бинарный файл в bin/hexlet-path-size
build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

# Устанавливает собранный бинарник в GOBIN, чтобы его можно было запускать из любого места.
install: build
	go install ./cmd/hexlet-path-size 

# Запуск линтера
lint:
	golangci-lint run

# Запуск тестов
test:
	go test -v ./...