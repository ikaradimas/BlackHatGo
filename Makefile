.PHONY: rude_scanner

.rude_scanner: 
	go build -o bin/rude_scanner cmd/rude_scanner/main.go

.echo_server:
	go build -o bin/echo_server cmd/echo_server/main.go

.tcp_server:
	go build -o bin/tcp_proxy cmd/tcp_proxy/main.go

all: .rude_scanner .echo_server .tcp_server
