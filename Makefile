BIN=$(GOPATH)/bin
GOLINT = $(BIN)/golangci-lint

lint: | $(BASE) $(GOLINT)
	$(GOLINT) run --tests=false