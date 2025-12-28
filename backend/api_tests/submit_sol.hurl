# Test the run code endpoint
# /code/submit/language_name/:problem_id
POST http://localhost:8080/code/submit/javascript/1
Content-Type: application/json

{
    "code": "function test(input) { return 'input' }"
}
HTTP 200

[Asserts]
jsonpath "$.message" == "all tests ran"