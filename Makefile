run:
	go run ./cmd/atproxy

test:
	go test -v ./... -coverprofile cover.out -covermode=atomic | tee tests.out
