test:
	go test ./... -cover

coverage:
	go test ./... -coverprofile=c.out && go tool cover -html=c.out
