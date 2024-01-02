
build:
	go mod tidy	
	GOOS=linux GOARCH=amd64 go build .

clean:
	git clean -fd

.PHONY: clean
