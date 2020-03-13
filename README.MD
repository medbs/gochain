### Prerequisites
* Go 1.13

### Run the project

Start the ledger
```
go run cmd/main.go bc --l 9090
```

Start http server
```
go run cmd/main.go hs --p 8099
```

Add data to the ledger

```
curl -d '{"field":"t1"}' -H "Content-Type: application/json" -X POST http://localhost:8099/api/v1/ledger
```

### Build & run the project

Build 
```
go build -o dist/chain cmd/main.go
```
Run 
```
dist/chain serve --l 9090
```

