# Receipt Processor Service 


### Introduction

A simple receipt processor API that takes in receipt data and calculates award points based on a set of defined business rules. 

This projects uses the lightweight web framrwork [Echo](https://echo.labstack.com/) to set up an API server. The API offers a `POST` endpoint to upload receipts and return unique ID. It also offers a `GET` endpoint that takes an ID and returns point totals representing points awarded. 

### How to run the API 
Pre-requisites: this project is built with `go 1.21.0` 
Assumption: the environment where this API will have golang installed in a docker container. 

#### Installation

Run `go get` to install packages. This will install the `Echo` web framework and its middleware as well as the UUID generator. 

#### Run

to run the API simply execute, `go run main.go`. This will start a web server on the default `Echo` port: 1323. 

#### Sample input and output. 
Interact with the API with an http requester, e.g., `curl`. 

- Sample POST request (note the Content-Type header): 

    ```
    > curl -v localhost:1323/receipts/process -d '{               "retailer": "M&M Corner Market",
    "purchaseDate": "2022-03-20",
    "purchaseTime": "14:33",
    "items": [
        {
        "shortDescription": "Gatorade",
        "price": "2.25"
        },{
        "shortDescription": "Gatorade",
        "price": "2.25"
        },{
        "shortDescription": "Gatorade",
        "price": "2.25"
        },{
        "shortDescription": "Gatorade",
        "price": "2.25"
        }
    ],
    "total": "9.00"
    }' -H 'Content-Type: application/json'
    ```
- Response (note the ID): 

    ```
    *   Trying 127.0.0.1:1323...
    * Connected to localhost (127.0.0.1) port 1323 (#0)
    > POST /receipts/process HTTP/1.1
    > Host: localhost:1323
    > User-Agent: curl/8.1.2
    > Accept: */*
    > Content-Type: application/json
    > Content-Length: 409
    >
    < HTTP/1.1 201 Created
    < Content-Type: application/json; charset=UTF-8
    < Date: Mon, 04 Sep 2023 14:44:54 GMT
    < Content-Length: 46
    <
    {"id":"1b24c44f-a9bb-4617-8cee-18e8462ead9a"}
    * Connection #0 to host localhost left intact
    ```

- Sample GET request: 
    ```
    > curl -v localhost:1323/receipts/ffe0f5d6-5944-4f9f-a825-2ae412700004/points
    ```
- Response: 
    ```
    *   Trying 127.0.0.1:1323...
    * Connected to localhost (127.0.0.1) port 1323 (#0)
    > GET /receipts/ffe0f5d6-5944-4f9f-a825-2ae412700004/points HTTP/1.1
    > Host: localhost:1323
    > User-Agent: curl/8.1.2
    > Accept: */*
    >
    < HTTP/1.1 200 OK
    < Content-Type: application/json; charset=UTF-8
    < Date: Mon, 04 Sep 2023 14:47:55 GMT
    < Content-Length: 15
    <
    {"points":109}
    * Connection #0 to host localhost left intact
    ```

---
### Code Structure 

The code is separated into packages similar to a traditional model-controller structure with a handler for http request handling.

- `model` package for data models. 

- `processor` package with for handling the business logic for how to calculate award points. Tests are available in `processor_test` as well. 

- `handler` package contains  for HTTP request handling.


#### Package Deep Dive

- `handler` package: this package contains the `Handler` struct which acts as an HTTP endpoint handler, with methods for POST receipts and GET a receipt's award points based on and ID. 

    - `ProcessReceipt` function takes in the echo Context, passes the POST request data to the `Processor` struct which calculates the points and returns a unique ID for that receipt for retrieval. 
    - `GetReceiptPoints` function takes in the echo Context and a path parameter `id` to locate a receipt and return the award points.

- `processor` package: This package contains the `Processor` struct which handles the business rules defined for calculating award points based on data in the receipt

#### Rationale and future improvement scenarios

Isolating models from data processing and HTTP handling into separate packages allows for better future modification and maintainability of code. 

Examples: 

`processor` changes 

- Introduce new rewards calcuations can be handled by simply creating a new function.

- Changing the rules for a specific calculation is isolated to helper functions. E.g., changing the calculation for retailer name is isolated in the `calcPointsRetailerName` helper function. 

- Data storage. A new interface can be injected in the `processor` struct, e.g., `"database/sql"` go package. That way, receipt data and award points can be persisted for better analytics. 

`model` changes 

- if a new data model is introduced, then a new struct can be added to the `model` package. 

`handler` changes
- For HTTP request handling, `handler` can be expanded to handle explicit HTTP errors.

- Other improvements may include authentication, proxying, etc. 

---

### Testing

Basic tests are included to check `processor` calculations for retailer name, dates and time, items, as well as receipt totals. 

To run tests, execute: `go test ./...`
