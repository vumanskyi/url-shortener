# Variables
TEST_DIR=./internal
COVERAGE_DIR=coverage
COVERAGE_FILE=$(COVERAGE_DIR)/coverage.out
COVERAGE_HTML=$(COVERAGE_DIR)/coverage.html

# Default target
all: test

# Run tests with code coverage (only for internal)
test: $(COVERAGE_FILE)

$(COVERAGE_FILE):
	@echo "Running tests in $(TEST_DIR)..."
	@mkdir -p $(COVERAGE_DIR)
	go test $(TEST_DIR)/... -coverprofile=$(COVERAGE_FILE)
	@echo "Tests completed."

# Display code coverage in the terminal (ensure it always runs after test)
coverage: $(COVERAGE_FILE)
	@echo "Showing code coverage..."
	go tool cover -func=$(COVERAGE_FILE)

# Generate HTML report and open it in the browser (ensure it always runs after test)
coverage-html: $(COVERAGE_FILE)
	@echo "Generating HTML coverage report..."
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "Opening coverage report in browser..."
	@open $(COVERAGE_HTML) || xdg-open $(COVERAGE_HTML)

# Clean up generated files
clean:
	@echo "Cleaning up..."
	rm -rf $(COVERAGE_DIR)
	@echo "Cleanup complete."

# Help
help:
	@echo "Makefile Commands:"
	@echo "  make test           - Run tests with coverage (for internal only)"
	@echo "  make coverage       - Display code coverage in the terminal"
	@echo "  make coverage-html  - Generate and open an HTML code coverage report"
	@echo "  make clean          - Remove generated files"
