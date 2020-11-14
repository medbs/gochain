### Prerequisites
* Go 1.13

### Run the project

Start the blockchain
```
go run cmd/main.go bc --address 127.0.0.1 --ptpPort 9090 --httpPort 8090
```

Add data to the blockchain

```
curl  -X POST http://localhost:8090/api/v1/ledger?data=helloworld
```

Read Data from the blockchain

```
curl -X GET http://localhost:8090/api/v1/ledger
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

