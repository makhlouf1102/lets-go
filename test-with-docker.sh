#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo "Running tests..."

# Run tests with verbose flag, race detector, and coverage report
go test -v -race -cover ./... | while read -r line; do
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
    go test ./... -coverprofile=./test_docs/coverage.out
    go tool cover -html=./test_docs/coverage.out -o ./test_docs/coverage.html
    
    echo -e "${GREEN}Coverage report generated as ./test_docs/coverage.html${NC}"
else
    echo -e "\n${RED}Tests failed!${NC}"
    exit $TEST_EXIT_CODE
fi 