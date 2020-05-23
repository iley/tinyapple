demo: always
	GOOS=linux GOARCH=mipsle go build -mod=vendor github.com/iley/tinyapple/cmd/demo

check: always
	GOOS=linux GOARCH=mipsle go build -mod=vendor github.com/iley/tinyapple/cmd/check

.PHONY: always
