run:
	go run main.go

kill:
	sudo kill $(sudo lsof -t -i:9898)
	
spec:
	cd spec && go generate
	