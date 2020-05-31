MIPSBIN=demo
HOSTBIN=

default: $(MIPSBIN) $(HOSTBIN)

$(MIPSBIN): %: always
	GOOS=linux GOARCH=mipsle go build -mod=vendor github.com/iley/tinyapple/cmd/$@

$(HOSTBIN): %: always
	go build -mod=vendor github.com/iley/tinyapple/cmd/$@

.PHONY: default always
