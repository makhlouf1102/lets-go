# Test the run code endpoint
POST http://localhost:8080/code/run
Content-Type: application/json

{
    "source_code": "console.log('Hello World')",
    "language_id": 63
}