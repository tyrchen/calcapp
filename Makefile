CHECK=\033[32m✔\033[39m
DONE="\n$(CHECK) Done.\n"

GO=go
BIN=./bin
DATA=/var/data/calcapp
PACKAGES=network utils calc
PACKAGE_PATHS=$(patsubst %,./%, $(PACKAGES))
TARGETS=$(patsubst %.go,$(BIN)/%,$(wildcard *.go))
ALL_FILES=$(shell find . -type f -name '*.go')

.PHONY: build packages clean dir

build: clean packages $(TARGETS)
	@echo $(DONE)

$(TARGETS): $(BIN)/%: %.go
	@echo "building $<..."
	@$(GO) build -o $@ $< 

packages:
	@echo "making packages"
	@$(GO) install $(PACKAGE_PATHS)


dir:
	@mkdir -p $(DATA)/bp/origin $(DATA)/bp/new $(DATA)/mp
	@mkdir -p $(DATA)/mac

clean:
	@rm -f $(TARGETS)
	@echo $(DONE)
