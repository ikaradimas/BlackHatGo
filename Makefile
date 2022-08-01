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

.negroni_example:
	go build -o bin/negroni_example cmd/negroni_example/main.go

.html_template_example:
	go build -o bin/html_template_example cmd/html_template_example/main.go

.credential_harvester:
	go build -o bin/credential_harvester cmd/credential_harvester/main.go

.websocket_keylogger:
	go build -o bin/websocket_keylogger cmd/websocket_keylogger/main.go

.serve:
	go build -o bin/serve cmd/serve/main.go

.fqdn:
	go build -o bin/fqdn cmd/fqdn/main.go

.guessdns:
	go build -o bin/guessdns cmd/guessdns/main.go

.simple_dns_server:
	go build -o bin/simple_dns_server cmd/simple_dns_server/main.go

.mongo_example:
	go build -o bin/mongo_example cmd/mongo_example/main.go	

.mysql_example:
	go build -o bin/mysql_example cmd/mysql_example/main.go	

.postgres_example:
	go build -o bin/postgres_example cmd/postgres_example/main.go	

.mongo_miner:
	go build -o bin/mongo_miner cmd/mongo_miner/main.go

.mysql_miner:
	go build -o bin/mysql_miner cmd/mysql_miner/main.go

.fs_scanner:
	go build -o bin/fs_scanner cmd/fs_scanner/main.go

.find_network_devices:
	go build -o bin/find_network_devices cmd/find_network_devices/main.go

.simple_bpf_filter:
	go build -o bin/simple_bpf_filter cmd/simple_bpf_filter/main.go

all: .rude_scanner \
	 .echo_server \
	 .tcp_server \
	 .command_proxy \
	 .shodan \
	 .metasploit \
	 .bing_scrape \
	 .simple_http_server \
	 .middleware_example \
	 .negroni_example \
	 .html_template_example \
	 .credential_harvester \
	 .websocket_keylogger \
	 .serve \
	 .fqdn \
	 .guessdns \
	 .simple_dns_server \
	 .mongo_example \
	 .mysql_example \
	 .postgres_example \
	 .mongo_miner \
	 .mysql_miner \
	 .fs_scanner \
	 .find_network_devices \
	 .simple_bpf_filter