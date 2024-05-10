# total-coder-w19

##### Given a coverage profile produced by 'go test':

```go
go test -coverprofile=c.out
```

##### Open a web browser displaying annotated source code

```go
go tool cover -html=c.out
```

**Write out an HTML file instead of launching a web browser:**

```sh
$ go tool cover -html=c.out -o coverage.html
```

