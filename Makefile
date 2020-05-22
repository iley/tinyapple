demo: always
	GOOS=linux GOARCH=mipsle go build -mod=vendor github.com/iley/tinyapple/cmd/demo

.PHONY: always
