all : rfk rfk-server rfk-cli rfk-ident

rfk :
	go build

rfk-server :
	go build -o rfk-server ./server

rfk-cli :
	go build -o rfk-cli ./cli

rfk-ident :
	go build -o rfk-ident ./ident

clean :
	rm ./rfk || true
	rm ./rfk-cli || true
	rm ./rfk-server || true
	rm ./rfk-ident || true
