generate:
	go get github.com/99designs/gqlgen@v0.17.49
	go generate ./...

run:
	go run ./server.go


generate_mock:
	go get github.com/golang/mock/mockgen
	go generate