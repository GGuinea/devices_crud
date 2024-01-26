build:
	go build -o bin/main main.go

run:
	go run main.go

test:
	go clean -testcache && go test -v ./...

clean:
	rm -rf bin

docker-build-run:
	docker run -p 0.0.0.0:8080:8080 $(docker build -q .)
