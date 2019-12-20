test:
	go test ./... -cover

coverage:
	go test ./... -coverprofile=c.out && go tool cover -html=c.out

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

build_windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -installsuffix cgo -o app.exe .

build_rpi:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -a -installsuffix cgo -o app .
