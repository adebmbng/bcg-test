# Description

Hi, you'll find the source codes of the BCG take home test here. This repo combines the DDD pattern and functional
programming. This pattern made by me as the result of my experience in golang in professional work and side-projects.

````
|-- bcg-test
    |-- domains
        |-- inventory
        |-- promotion
    |-- configs
    |-- migrations
    |-- entities
    |-- handler
    |-- pkg
    |-- repositories
````

Entities folder is where the business logic are.

# How to Run

## Initial

Run `go mod vendor` to download everything to vendor folder

## Setup DB

Run every query on migrations folder

## Unit test

This repo has 3 unit test, for test scenario it is at `domains/inventories/inventory_test`
function `calculateFinalPrice`. But, take a look at `promotion_test.go` I have something cool right there, full unit
test for each scenario.

## GraphQL

My time is running out, I need to do something urgent. So, my plan was to add graphql server on `main.go` and call
function on `handler` then the handler will call the `inventory domain`

# What I miss in this Repo
* Validating stock on query
* GraphQL
* I have 3 TODO, 2 of them is because of the limitation of the test case. (search `TODO`)


