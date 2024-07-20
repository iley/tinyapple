MIPSBIN=demo bot clock
HOSTBIN=

SSH_HOST=littlemac

default: $(MIPSBIN) $(HOSTBIN)

$(MIPSBIN): %: always
	GOOS=linux GOARCH=mipsle go build -mod=vendor github.com/iley/littlemac/cmd/$@

$(HOSTBIN): %: always
	go build -mod=vendor github.com/iley/littlemac/cmd/$@

scp: clock
	scp clock $(SSH_HOST):/root

deploy: clock
	ssh $(SSH_HOST) /etc/init.d/clock stop
	sleep 5
	scp clock $(SSH_HOST):/root
	ssh $(SSH_HOST) /etc/init.d/clock start

internal/fonts/digits.go: digits.bdf
	tinyfontgen --package fonts --fontname Digits $^ --output $@ --all --verbose

.PHONY: default always deploy
