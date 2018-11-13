server:
	GOOS=linux GOARCH=amd64 go build -o gates-server ./server

.PHONY: server

serve:
	open http://127.0.0.1:3000
	php -S 127.0.0.1:3000

.PHONY: serve
