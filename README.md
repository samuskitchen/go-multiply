# Multiply Api
This is a project for a challenge in Golang. This api is responsible for multiplying two values

The following technologies were used in this project:
- [Golang 1.15](https://golang.org/dl/)
- [go-chi](https://github.com/go-chi/chi)
- [go-chi/cors](https://github.com/go-chi/cors)
- [objx](https://github.com/stretchr/objx)
- [testify](https://github.com/stretchr/testify)
- [mockery](https://github.com/vektra/mockery)
- [yalm](https://github.com/go-yaml/yaml/tree/v3)
- [healt](https://github.com/dimiro1/health)

## Generate Mock Interface
This is an automatic mock generator using mockery, the first thing we must do is go to the path of the file that we want to autogenerate:

Download the library
```
go get github.com/vektra/mockery/v2/.../
```

We enter the route where you are
```
cd path
```

After entering the route we must execute the following command, Repository this is name the interface
```
mockery --name=Repository
```

## Test commands for the project
These are the commands to run the unit and integration tests of the project

#### Execute All Test
This is the command to run the white box tests, and the test report command
```
go test -v -coverprofile=coverage.out -coverpkg=./domain/... ./test/...

go tool cover -html=coverage.out
```

This command gets the total coverage of the project
```
go tool cover -func coverage.out
```

#### Execute All Test Integration
This is the command to run the black box tests, and the test report command
```
go test -v -coverprofile=coverage_integration.out -coverpkg=./domain/... ./test/integration

go tool cover -html=coverage_integration.out
```

This command gets the total coverage of test integration
```
go tool cover -func coverage_integration.out
```