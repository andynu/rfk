all : server cli ident

server :
	go build -o rfk-server github.com/andynu/rfk/server
cli :
	go build -o rfk-cli github.com/andynu/rfk/cli
ident :
	go build -o rfk-ident github.com/andynu/rfk/ident
