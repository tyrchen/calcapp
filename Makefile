CHECK=\033[32m✔\033[39m
DONE="\n$(CHECK) Done.\n"

GO=go
BIN=./bin
TARGETS=$(patsubst %.go,$(BIN)/%,$(wildcard *.go))

.PHONY: build clean dir

build: $(TARGETS)
	@echo $(DONE)

$(TARGETS): $(BIN)/%: %.go
	@echo "building $<..."
	@$(GO) build -o $@ $< 

dir:
	@mkdir -p $(BIN)/data/bp
	@mkdir -p $(BIN)/data/mac

clean:
	@rm -f $(TARGETS)
	@echo $(DONE)
