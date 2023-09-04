# Receipt Processor Service 

A simple receipt processor API that takes in receipt data and calculates award points based on a set of defined business rules. See structure of the API below.

### How to run the API 
Pre-requisites: this project is built with `go 1.21.0` 
Assumption: the environment where this API will have golang installed in a docker container. 

#### Installation
There are some go packages that are needed to run this API. 
- The API is built using the lightweight web framrwork Echo

#### Run

to run the API simply execute, `go run main.go`


---
### Code Structure 

#### Packages
- `handler` package: this package contains the `Handler` struct which acts as an HTTP endpoint handler, with methods for POST receipts and GET a receipt's points based on and ID. 
    - `ProcessReceipt` function takes in the echo Context, passes the POST request data to the `Processor` struct which calculates the points and returns a unique ID for that receipt for retrieval. 
    - `GetReceiptPoints` function takes in the echo Context and a path parameter `id` to locate a receipt and return the award points.

- `processor` package: This package contains the `Processor` struct which handles the business rules defined for calculating award points based on data in the receipt

...

// TBD 
#### Models 

// TBD
---

### Testing

// TBD
