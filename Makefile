# Определяем переменную для команды go
GO_CMD=go

# Определяем переменную для пути к вашему приложению
APP_PATH=./cmd/app/main.go

# Цель по умолчанию для запуска приложения
run:
	$(GO_CMD) run $(APP_PATH)

# Цель для сборки приложения
build:
	$(GO_CMD) build -o ./bin/app $(APP_PATH)

# Цель для очистки собранных файлов
clean:
	rm -f ./bin/app

# Цель для запуска тестов
test:
	$(GO_CMD) test ./...

.PHONY: run build clean test
