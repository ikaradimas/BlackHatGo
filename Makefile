.PHONY: rude_scanner

.rude_scanner: 
	go build -o bin/rude_scanner cmd/rude_scanner/main.go

all: .rude_scanner
