@rem go build -race -o ./bin/server.exe ./test/server/main.go
go build -o ./bin/server.exe ./test/server/main.go
go build -o ./bin/client.exe ./test/client/main.go

@pause