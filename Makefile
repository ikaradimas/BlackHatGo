.rude_scanner: 
	go build -o bin/rude_scanner cmd/rude_scanner/main.go

.echo_server:
	go build -o bin/echo_server cmd/echo_server/main.go

.tcp_server:
	go build -o bin/tcp_proxy cmd/tcp_proxy/main.go

.command_proxy:
	go build -o bin/command_proxy cmd/command_proxy/main.go

.shodan:
	go build -o bin/shodan cmd/shodan/main.go

.metasploit:
	go build -o bin/metasploit cmd/metasploit/main.go

.bing_scrape:
	go build -o bin/bing_scrape cmd/bing_scrape/main.go

.simple_http_server:
	go build -o bin/simple_http_server cmd/simple_http_server/main.go

.middleware_example:
	go build -o bin/middleware_example cmd/middleware_example/main.go

all: .rude_scanner \
	 .echo_server \
	 .tcp_server \
	 .command_proxy \
	 .shodan \
	 .metasploit \
	 .bing_scrape \
	 .simple_http_server \
	 .middleware_example
	 
