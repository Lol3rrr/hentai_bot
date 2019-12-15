test:
	go test ./... -cover

coverage:
	go test ./... -coverprofile=c.out && go tool cover -html=c.out

coverage_total:
	go test ./... -coverprofile=c.out | go tool cover -func c.out | grep total | awk '{print $3}'
