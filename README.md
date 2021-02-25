## GO

- First you will need to [download and install go](https://golang.org/doc/install)
- The workspace directory for the GO project (GO_excersie) is as follows\
  at the root of the hierarchy:\
  bin/\
  src/\
  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;github.com/\
  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ahmedsheha4/\
  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;GO_excersie/\
  pkg/
- The external modules used for this assignment are [Sqlite3 driver](https://github.com/mattn/go-sqlite3)\
  This package can be installed with the go get command

```bash
go get github.com/mattn/go-sqlite3
```

- [Gorilla mux](https://github.com/gorilla/mux) for HTTP requests routing and handling
  This package can be installed with the go get command

```bash
go get -u github.com/gorilla/mux
```

- Set the environment variable GOPATH to the workspace directory
- Make sure the GO111MODULE environment variable is set to "auto"
- Run the following command to install all the missing dependencies

```bash
go install
```

- To run the main package

```bash
go run main.go
```

- The server will run on port 8000
- To automate the execution of test files and functions in the Go project use the following command

```bash
go test
```
