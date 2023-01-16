#!/usr/bin/make -f

test:
	@go version
	go fmt ./...
	go mod tidy
	@echo
	go test -cover -count=1 -timeout=1s -race ./...

doc:
	printf '# ' > README.md && \
		head -n 1 go.mod | sed 's/^module //' >> README.md && \
		echo >> README.md && \
		echo >> README.md && \
		go doc -all github.com/mdwhatcott/go-set/set  | sed 's/^/\t/' >> README.md

onefile:
	@go-mergepkg -dirs "set" -header "github.com/mdwhatcott/go-set@$(shell git describe) (a little copy-paste is better than a little dependency)"
