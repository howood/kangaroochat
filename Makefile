.PHONY: install,run,test,testv

install:
	cd /go/src/github.com/howood/kangaroochat/kangaroochat && export GO111MODULE=on && go install

run:
	export GO111MODULE=on && go run ./kangaroochat/kangaroochat.go -v

test:
	export GO111MODULE=on && go test ./...

testv:
	export GO111MODULE=on && go test ./... -v
