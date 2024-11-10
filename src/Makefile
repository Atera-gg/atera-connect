APP_NAME = atera-connect
BUILD_DIR = build

.PHONY: build run clean

build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	FYNE_RENDER=software GOOS=darwin GOARCH=arm64 go build -v -o $(BUILD_DIR)/$(APP_NAME) main.go
	@echo "Build complete: $(BUILD_DIR)/$(APP_NAME)"

run:
	@echo "Running $(APP_NAME)..."
	./$(BUILD_DIR)/$(APP_NAME)

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Cleanup complete"
