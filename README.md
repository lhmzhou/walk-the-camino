# walk-the-camino

`walk-the-camino` uses the Godog testing framework for writing a resuable BDD-based regression suite. Similar to any Godog application, `walk-the-camino` searches for test case specifications in `feature/*.feature` files found in the `tests/functional` folder, and then pairs them with the corresponding `FeatureABC(suite * godog.Suite)` functions. The backbone of the application was recycled using [have-some-func](https://github.com/lhmzhou/have-some-func) as the backend service.

## Core Structure

```
walk-the-camino
  ├── app
  │    ├── handlers_test.go
  │    ├── handlers.go
  │    ├── start.go
  │    └── package.json
  │
  ├── cert
  │    └── ...
  │
  ├── data
  │    └── employee.go
  │
  ├── database
  │    └── processor.go
  │
  ├── tests
  │    └── functional
  │         ├── reports
  │         │    ├── node_modules
  │         │    │    ├── ...
  │         │    │    └── cucumber-html-reporter
  │         │    └── main.js
  │         ├── testutil
  │         │   └── testutil.go   
  │         ├── Update-Post-Request
  │         │    ├── config
  │         │    ├── features
  │         │    ├── report
  │         │    ├── main.go
  │         │    ├── put_request.go
  │         │    ├── main_test.go
  │         │    └── report_test.go
  │         ├── cleanJson.go
  │         └── package.json
  │
  ├── utils
  │    └── constants.go
  │
  ├── vendor
  │    └── github.com
  │         ├── buger
  │         ├── DATA-DOG
  │         └── gorilla        
  │
  ├── ...
  ├── dev-compose.yml
  ├── Dockerfile
  ├── LICENSE 
  ├── project.go
  └── README.md
```

## Prerequisites

```
Go: go1.11.5 (protip: make sure Go binaries are installed by default)
godog: v0.7.9 
```

## Usage

Install following libraries:
1. go get github.com/cucumber/godog/cmd/godog@v0.9.0
2. go get -u github.com/gorilla/mux
3. go get -u https://github.com/buger/jsonparser

Build application and functional test cases
```
$ docker-compose -f dev-compose.yml up --exit-code-from functional --build functional
```

### Expected Output

On the terminal:
```
run PASS ok walk-the-camino/tests/functional/Update-Post-Request 1.000s [no tests to run]
```

On the UI:
<img width="1143" alt="ui screenshot" src="https://user-images.githubusercontent.com/16420802/79627252-b719e000-8104-11ea-959c-1336ae45f0df.png">


## Contributing

Pull requests are welcomed. Feel free to fork and submit a PR to fix a potential bug, suggest improvements, add new features, etc. Any breaking changes will be published in the [CHANGELOG](https://github.com/lhmzhou/walk-the-camino/blob/master/CHANGELOG.md).


## License 

`walk-the-camino` is licensed under [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/lhmzhou/walk-the-camino/blob/master/LICENSE)


## Additional Geekery

[Configure Go toolchain](https://golang.org/doc/install#testing)
</br>
[Cucumber Docs](https://cucumber.netlify.app/docs/installation/golang/)
</br>
[BDD Demo in Go](https://dev.to/jankaritech/demonstrating-bdd-behavior-driven-development-in-go-1eci)