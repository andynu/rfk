all : rfk-server rfk-cli rfk-ident

rfk-server :
	go build -o rfk-server ./server

rfk-cli :
	go build -o rfk-cli ./cli

rfk-ident :
	go build -o rfk-ident ./ident

clean :
	rm ./rfk || echo already cleaned.
	rm ./rfk-cli || echo already cleaned.
	rm ./rfk-server || echo already cleaned.
	rm ./rfk-ident || echo already cleaned.
