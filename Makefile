all:
	env GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o ./target/simple-api

clean:
	rm -rf ./target/simple-api
