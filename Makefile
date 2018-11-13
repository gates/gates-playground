server:
	GOOS=linux GOARCH=amd64 go build -o gates-server ./server

.PHONY: server
