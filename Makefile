tinyapple: main.go
	GOOS=linux GOARCH=mipsle go build -mod=vendor

clean:
	rm -f tinyapple

.PHONY: clean
