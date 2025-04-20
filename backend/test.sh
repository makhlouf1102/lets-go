#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Folders to exclude from testing
EXCLUDE_FOLDERS=("dockerController" "dockerClient")

# Convert excluded folders to a regex pattern for skipping
EXCLUDE_PATTERN=$(printf "|%s" "${EXCLUDE_FOLDERS[@]}")
EXCLUDE_PATTERN=${EXCLUDE_PATTERN:1} # Remove leading "|"

echo "Running tests, excluding folders: ${EXCLUDE_FOLDERS[*]}..."

# Run tests with verbose flag, race detector, and coverage report, skipping excluded folders
go test -v -race -cover $(go list ./... | grep -vE "$EXCLUDE_PATTERN") | while read -r line; do
    if [[ $line == *PASS* ]]; then
        echo -e "${GREEN}$line${NC}"
    elif [[ $line == *FAIL* ]]; then
        echo -e "${RED}$line${NC}"
    else
        echo "$line"
    fi
done

# Capture the exit code of the tests
TEST_EXIT_CODE=${PIPESTATUS[0]}

# Create coverage report if tests pass
if [ $TEST_EXIT_CODE -eq 0 ]; then
    echo -e "\n${GREEN}All tests passed!${NC}"
    
    # Generate coverage report
    echo -e "\nGenerating coverage report..."
    go test $(go list ./... | grep -vE "$EXCLUDE_PATTERN") -coverprofile=./test_docs/coverage.out
    go tool cover -html=./test_docs/coverage.out -o ./test_docs/coverage.html
    
    echo -e "${GREEN}Coverage report generated as ./test_docs/coverage.html${NC}"
else
    echo -e "\n${RED}Tests failed!${NC}"
    exit $TEST_EXIT_CODE
fi
