build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o env-info

clean:
	go clean