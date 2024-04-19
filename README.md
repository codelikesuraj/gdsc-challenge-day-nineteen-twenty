# Days 19-20: Write Unit Tests for Your Backend

Write unit tests for your backend. These unit tests should ensure that the backend is running smoothly and all routes return the expected results and response codes under defined conditions.

## Setup
- Navigate to the root of this repo.
- Run the command ```go run ./main.go``` to start the server.
- Visit the following url endpoints:
    |METHOD|DESCRIPTION|ENDPOINT|SAMPLE BODY|
    |------|-----------|--------|----|
    |POST  |Register user|http://127.0.0.1:3000/register|{"username":"username","password":"password"}|
    |POST  |Login user   |http://127.0.0.1:3000/login|{"username":"username","password":"password"}|
    |POST  |Refresh token   |http://127.0.0.1:3000/login|{"refresh_token": "refresh_token"}|
    |GET   |Get all books added by logged-in user|http://127.0.0.1:3000/books|-|
    |GET   |Get a book added by logged-in user|http://127.0.0.1:3000/books/{id}|-|
    |POST  |Create a book for logged-in user|http://127.0.0.1:3000/books|{"author":"book_author","title":"book_title"}|
- RUN TESTS with the command `go test ./...` or run with the verbose flag `go test ./... -v`
