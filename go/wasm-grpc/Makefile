web:
	cp `go env GOROOT`/misc/wasm/wasm_exec.js web/

proto: web
	protoc --proto_path=proto --go_out=plugins=grpc:proto echo.proto

client: proto
	GOARCH=wasm GOOS=js go build -o web/client.wasm client/client.go


.PHONY: proto client web

serve: client
	go run server/server.go
