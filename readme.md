# Supermarket API User Stories ![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

This API is responsible for add, delete, and fetch all produce in the system..

## Index
- [Run it locally](#run-it-locally)
- [Test the application](#test-the-application)

## Run it locally
To run the api locally it is necessary to clone this repository from Github:
````
git clone https://github.com/Jovanny159/project_1.git
````

After that, you have to move to the root of the project and run this command:
````
go run cmd/main.go
````

Now you have the API running in the port :8080 so you can invoke the endpoints with the following curl (you can use Postman too):

### POST Add Produce
You can add new produces with the following endpoit

````bash
curl --location --request POST 'http://localhost:8080/produce' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Lettuce",
    "code": "A12T-4GH7-QPL9-3N4M",
    "unit_price": 3.46
}'
`````

- The produce includes name, produce code, and unit price
- The produce name is alphanumeric and case insensitive
- The produce codes are sixteen characters long, with dashes - separating each four character group
- The produce codes are alphanumeric and case insensitive
- The produce unit price is a number with up to 2 decimal places

It will respond in the following way:
````JSON
{
    "name": "Lettuce",
    "code": "A12T-4GH7-QPL9-3N4M",
    "unit_price": 3.46
}
````

In case of error it will respond with the right status code and the following response:
````JSON
{
    "message": "error while adding the produce",
    "error": "internal_server_error",
    "status": 500,
    "cause": [
        "an error of type: bad_request with value: the provided data is invalid and cause: []"
    ]
}
````

### GET Produce
You can fetch all the produce in the database. 

````bash
curl --location --request GET 'http://localhost:8080/produce'
`````

It will respond in the following way with all the produce:
````JSON
[
    {
        "name": "Lettuce",
        "code": "A12T-4GH7-QPL9-3N4M",
        "unit_price": 3.46
    },
    {
        "name": "Lettuce",
        "code": "A12T-4GH7-QPL9-3N4N",
        "unit_price": 3.46
    }
]
````

In case of error it will respond with the right status code and the following response:
````JSON
{
    "message": "data not found",
    "error": "status_not_found",
    "status": 404,
    "cause": []
}
````

### DELETE Produce
Delete a produce item from the database. Accepts a url parameter of Produce Code that will identify the item to delete:

````bash
curl --location --request DELETE 'http://localhost:8080/produce/A12T-4GH7-QPL9-3N4M'
`````

It will respond with a 200 - StatusOK

In case of error it will respond with the right status code and the following response:
````JSON
{
    "message": "error",
    "error": "internal_server_error",
    "status": 500,
    "cause": []
}
````

Besides that, you have another endpoint to check the API health status:
````bash
curl --location --request GET 'http://localhost:8080/ping'
````

If the API is running successfully, it has to return the word ```pong``` 

## Test the application
This API have unit tests to warranties the integrity of the application.
You can run this tests with the following command from the root of the project:
````
go test ./...
````

#### Coverage
You can also generate a coverage report with the following commands:
- Generating the coverage.out file
````
go test -covermode=set -coverprofile=coverage.out ./...
````
- Displaying an HTML coverage report
 
````
 go tool cover -html=coverage.out
````