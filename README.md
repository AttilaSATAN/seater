# seater
Simple show case app

idiotically simple solution with a ***BONUS***

I didn't want to just calculate a 60x10 plane seats but I created a `plane` package that could define new planes and query their seats with their ids individually. For this check out the `/cmd/seater/main.go`

Dependecy management provided by `dep` you should [install](https://github.com/golang/dep/) it.
uses only `github.com/kataras/iris` though

```
go get github.com/attilasatan/seater
cd $GOPATH/src/github.com/attilasatan/seater
```
or clone this repo and in main folder. 
```
dep ensure
go run cmd/seater/main.go
```

There is only one end-point which gets the seat number as asked in the assignment docs.

localhost:8080/plane/id/{id:string}/seat/{index:int}

### TODO

- Tests
- Docs
