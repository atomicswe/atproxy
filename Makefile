run:
	ATPROXY_CONFIG_PATH=./config.json go run ./cmd/atproxy

test:
	go test -v ./... -coverprofile cover.out -covermode=atomic | tee tests.out

# needs sudo to create the directory where
# the executable will be placed (/opt/atproxy)
build:
	./build.sh
