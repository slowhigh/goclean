.PHONY:

run:
	go run ./cmd/server .; \

di:
	wire ./cmd/server; \

format:
	goimports -w .; \
	gofmt -s -w .; \

docs:
	swag init -pd -g ./cmd/server/main.go -o ./third_party/docs/; \












