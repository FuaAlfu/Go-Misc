---
stack: Go, Rest
lang: all, Go 1.22
---

# Building Rest API in Go 1.22
using new features..

## run program
```
go run main.go
```

## run local host
```
curl http://localhost:8080
```

## return all comments
```
curl http://localhost:8080/comment
```

## return a single comment
```
curl http://localhost:8080/comment/id
```

## post a new comment
```
curl -X POST http://localhost:8080/comment
```