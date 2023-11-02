rungateway:
	cd gateway && cd cmd && go run main.go
runaccount:
	cd account && cd cmd && go run main.go
runcrypto:
	cd account && cd cmd && go run main.go

run:
	runaccount
	runcrypto
	rungateway