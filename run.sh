# Define the binary name based on OS
if [ "$OSTYPE" == "msys" ] || [ "$OS" == "Windows_NT" ]; then
    BINARY_NAME="lets-go.exe"
else
    BINARY_NAME="lets-go"
fi

# Clean up existing binary
if [ -f "$BINARY_NAME" ]; then
    rm "$BINARY_NAME" || exit 1
fi

# Build the application
echo "Building application..."
go build -o ./build/"$BINARY_NAME" || exit 1

# Run the application
echo "Starting application..."
./build/"$BINARY_NAME"